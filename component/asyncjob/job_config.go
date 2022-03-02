package asyncjob

import "time"

type jobConfig struct {
	Name       string
	Retries    []time.Duration
	MaxTimeout time.Duration
}

var (
	defaultRetryTimes = []time.Duration{time.Second, time.Second * 2, time.Second * 5}
	defaultMaxTimeout = time.Minute * 2
)
