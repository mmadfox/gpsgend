package broker

import "time"

type Options struct {
	HistoryQueueCapacity int
	HistoryEnable        bool
	HistoryTimePeriod    time.Duration
}

func DefaultOptions() *Options {
	return &Options{
		HistoryQueueCapacity: 64,
		HistoryEnable:        true,
		HistoryTimePeriod:    time.Minute,
	}
}

func (o *Options) sanitize() {
	if o.HistoryQueueCapacity < 3 {
		o.HistoryQueueCapacity = 3
	}
	if time.Duration(o.HistoryTimePeriod.Seconds()) < 10 {
		o.HistoryTimePeriod = 10 * time.Second
	}
}
