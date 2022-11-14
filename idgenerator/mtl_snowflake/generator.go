package mtl_snowflake

import (
	"context"
	"errors"
	"fmt"
	"github.com/windeal/go-sdk/idgenerator/logger"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Generator struct {
	mutex   *sync.Mutex //互斥锁，保证线程安全
	config  *Config     // 可配置参数
	preSets *PreSets    // 基于config计算的参数值

	timelineProgress []int64 //各时间线进度
	curTimeline      int64   //当前时间线
	seq              uint64  //当前序号
}

type Config struct {
	TimeBit      uint64 // 时间位长度
	MachineIDBit uint64 // 实例ID位长度
	TimelineBit  uint64 // 时间线位长度
	SeqBit       uint64 // 序号位长度
	Epoch        int64  // 时间位的基准时间(unix nano)
	machineID    int64  // 机器码

}

type PreSets struct {
	// **** 偏移 ****
	shiftTimeBit      uint64 // 时间位偏移, 用于位计算得到ID中的时间部分
	shiftMachineIDBit uint64 // 机器码位偏移, 用于位计算得到ID中的机器码部分
	shiftTimelineBit  uint64 // 时间线偏移, 用于位计算得到ID中的时间线部分
	shiftSeq          uint64 // 序列号偏移, 用于位计算得到ID中的序列号部分

	// **** 掩码 ****
	maskTime      uint64
	maskMachineID uint64
	maskTimeline  uint64
	maskSeq       uint64 // 时间位偏移, 用于位计算得到ID中的时间部分

	// **** 最大值 ****
	maxTime      uint64 // 时间值的上限
	maxMachineID uint64 // 机器码值的上限
	maxTimeline  uint64 // 时间线的上限
	maxSeq       uint64 // 序列号值的上限
}

// NewGenerator : 新建一个生成器实例
func NewGenerator(ctx context.Context, opts ...Option) (gen *Generator, err error) {

	// 实例化Generator, 并配置初始值
	gen = &Generator{
		mutex: new(sync.Mutex),
		config: &Config{
			TimeBit:      defaultTimeBit,
			MachineIDBit: defaultMachineIDBit,
			TimelineBit:  defaultTimelineBit,
			SeqBit:       defaultSeqBit,
			Epoch:        defaultEpoch,
			machineID:    -1,
		},
		preSets:          nil,
		timelineProgress: make([]int64, 16), // 时间线不要超过16条
		curTimeline:      0,
		seq:              0,
	}

	machineID, _ := getIPSuffix(ctx)
	gen.config.machineID = int64(machineID)

	for _, opt := range opts {
		opt.apply(gen)
	}

	// 计算预设参数
	err = gen.buildPreSets(ctx)
	if err != nil {
		logger.LogErrorContextf(ctx, "gen.buildPreSets error, %+v")
		return gen, err
	}
	logger.LogInfoContextf(ctx, "gen.buildPreSets success")

	// sanity check
	err = gen.sanityCheck(ctx)
	if err != nil {
		logger.LogErrorContextf(ctx, "gen.sanityCheck error, %+v")
		return gen, err
	}
	logger.LogInfoContextf(ctx, "gen.sanityCheck success")

	return gen, err
}

func (gen *Generator) Generate(ctx context.Context) (int64, error) {
	gen.mutex.Lock()
	defer gen.mutex.Unlock()

	//settings := idGen.settings
	curTime := gen.unixNanoToOffsetTime(time.Now().UnixNano())
	if curTime < 0 {
		return 0, errors.New("epoch time is later than current time")
	}

	progress := gen.timelineProgress[gen.curTimeline] //当前时间线进度
	// **** 处理时钟回退 ****
	if curTime < progress { // 当前时间小于时间进度，说明发生了回退. 要么原地等待，要么切换一条时间线
		if progress-curTime < maxWaitTime { // 小幅回退，可以等待一会儿
			time.Sleep(time.Duration((progress - curTime) * timeUnit))
			curTime = gen.unixNanoToOffsetTime(time.Now().UnixNano())
		} else {
			// 查找合适的时间线
			timeline, err := gen.findSuitableTimeLine(curTime)
			if err != nil {
				return 0, err
			}

			// 切换时间线
			gen.curTimeline = timeline                // 切换时间线
			progress = gen.timelineProgress[timeline] // 更新时间线进度
			logger.LogInfoContextf(ctx, "timeline changed to %d", gen.curTimeline)
		}
	}

	if curTime == progress { // 当前时间与时间进度一致；如果序列号还有剩余，则消耗序列号，否则等待下一个时间单位
		gen.seq++ // 下一个序列号
		gen.seq = gen.seq & gen.preSets.maskSeq

		//如果当前时间单位的序号已用完，等待直到下一个时间单位
		if gen.seq == 0 {
			time.Sleep(time.Duration(timeUnit)) // 等待一个时间单位
			curTime = gen.unixNanoToOffsetTime(time.Now().UnixNano())
		}
	}

	if curTime > progress { // 当前时间已经晚于时间线进度
		gen.timelineProgress[gen.curTimeline] = curTime // 推进时间线进度
		gen.seq = 0
	}

	// 时间耗尽
	if uint64(curTime) > gen.preSets.maxTime {
		return 0, errors.New("out of time limit, maybe should reset epoch")
	}

	// 组装ID
	id := (curTime << gen.preSets.shiftTimeBit) |
		(gen.config.machineID << gen.preSets.shiftMachineIDBit) |
		(gen.curTimeline << gen.preSets.shiftTimelineBit) |
		int64(gen.seq)
	return id, nil
}

func (gen *Generator) buildPreSets(ctx context.Context) (err error) {
	gen.preSets = &PreSets{}

	// ID各个组成部分的偏移量
	gen.preSets.shiftTimeBit = gen.config.SeqBit + gen.config.TimelineBit + gen.config.MachineIDBit
	gen.preSets.shiftMachineIDBit = gen.config.SeqBit + gen.config.TimelineBit
	gen.preSets.shiftTimelineBit = gen.config.SeqBit
	gen.preSets.shiftSeq = 0

	// ID 各个组成部分最大值
	gen.preSets.maxSeq = (1 << gen.config.SeqBit) - 1
	gen.preSets.maxTimeline = (1 << gen.config.TimelineBit) - 1
	gen.preSets.maxMachineID = (1 << gen.config.MachineIDBit) - 1
	gen.preSets.maxTime = (1 << gen.config.TimeBit) - 1

	// ID 各个组成部分掩码
	gen.preSets.maskSeq = ((1 << gen.config.SeqBit) - 1) << gen.preSets.shiftSeq
	gen.preSets.maskTimeline = ((1 << gen.config.TimelineBit) - 1) << gen.preSets.shiftTimelineBit
	gen.preSets.maskMachineID = ((1 << gen.config.MachineIDBit) - 1) << gen.preSets.shiftMachineIDBit
	gen.preSets.maskTime = ((1 << gen.config.TimeBit) - 1) << gen.preSets.shiftTimeBit

	return nil
}

func (gen *Generator) sanityCheck(ctx context.Context) (err error) {

	// 最终生成的分布式唯一ID为int64， 第一位是符号位， 其他部分加起来需要为63位
	if 63 != gen.config.TimeBit+gen.config.MachineIDBit+gen.config.TimelineBit+gen.config.SeqBit {
		return errors.New("TimeBit+MachineIDBit+TimelineBit+SeqBit !=63")
	}

	curTime := gen.unixNanoToOffsetTime(time.Now().UnixNano())
	// 校验基准时间，基准时间不能晚于当前时间
	if curTime < 0 {
		return errors.New("epoch time must not later than current time")
	}
	// 当前时间和基准时间的偏移量已经超过了ID中时间位部分能存储的最大值了，一般是epoch设置的时间太早了
	if uint64(curTime) > gen.preSets.maxTime {
		return errors.New("the offset between the current time and the epoch time exceeds the limit, maybe epoch time set too early")
	}

	// 时间线最多16条，占用位数不要超过4bit
	if gen.config.TimelineBit > 4 {
		return errors.New("TimelineBit too big, not bigger than 4")
	}

	// 校验机器码是否合法
	if gen.config.machineID < 0 || gen.config.machineID > int64(gen.preSets.maxMachineID) {
		return errors.New(fmt.Sprintf("machineID must in [0, %d]", gen.preSets.maxMachineID))
	}

	return nil
}

func (gen *Generator) unixNanoToOffsetTime(unixNano int64) int64 {
	return (unixNano - gen.config.Epoch) / timeUnit
}

// findSuitableTimeLine 查找满足当前时间要求的时间线
func (gen *Generator) findSuitableTimeLine(curTime int64) (int64, error) {
	var maxProgress int64 = -1
	var timelineFound int64 = -1

	// 时间线进度必须小于curTime
	// 优先使用进度最快的时间线，这样进度慢的时间线可以用于支持更大的时间回退
	for index, progress := range gen.timelineProgress {
		if progress < curTime && progress > maxProgress {
			maxProgress = progress
			timelineFound = int64(index)
		}
	}
	if timelineFound == -1 {
		return -1, errors.New("cannot found suitable timeline")
	}

	return timelineFound, nil
}

// getIPSuffix : 轮询网卡，取第一个非本地回环的IPv4地址的最后一段，如192.168.0.122, 取122
// 如果取不到，返回-1
func getIPSuffix(ctx context.Context) (int, error) {
	addressList, err := net.InterfaceAddrs()
	if err != nil {
		logger.LogErrorContextf(ctx, "net.InterfaceAddrs error, %+v", err)
		return -1, err
	}
	ipStr := ""
	for _, address := range addressList {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ipStr = ipNet.IP.String()
				break
			}
		}
	}
	if len(ipStr) == 0 {
		logger.LogErrorContextf(ctx, "no ip found")
		return -1, err
	}
	ipArr := strings.Split(ipStr, ".")
	suffix, err := strconv.Atoi(ipArr[3]) //节点id
	if err != nil {
		logger.LogErrorContextf(ctx, "parse ip addr error, %+v", err)
		return -1, err
	}
	return suffix, err
}
