package mtl_snowflake

import "time"

const (
	defaultTimeBit      uint64 = 41                                                             //时间位数(可使用64年)
	defaultMachineIDBit uint64 = 9                                                              //实例ID位数(512实例)
	defaultTimelineBit  uint64 = 1                                                              //时间线位数,处理时钟回退
	defaultSeqBit       uint64 = 63 - defaultTimeBit - defaultMachineIDBit - defaultTimelineBit //序号位数
	timeUnit            int64  = 1e6                                                            //时间单位(1e6相当于ms)
	maxWaitTime         int64  = 1                                                              //当时间出现小幅回退时(这里设置为1时间单位)，等待时间递进到回退前时间再继续
)

var (
	defaultEpoch int64 = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()
)
