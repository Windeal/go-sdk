package mtl_snowflake

import "time"

const (
	defaultTimeBit      uint64 = 40                                                             //时间位数(可使用30年)
	defaultMachineIDBit uint64 = 15                                                             //实例ID位数(32768实例)
	defaultTimelineBit  uint64 = 2                                                              //时间线位数,处理时钟回退
	defaultSeqBit       uint64 = 63 - defaultTimeBit - defaultMachineIDBit - defaultTimelineBit //序号位数, 默认6bit，即每毫秒生成2^6=64个ID
	timeUnit            int64  = 1e6                                                            //时间单位(1e6相当于ms)
	maxWaitTime         int64  = 1                                                              //当时间出现小幅回退时(这里设置为1时间单位)，等待时间递进到回退前时间再继续
)

var (
	defaultEpoch int64 = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()
)
