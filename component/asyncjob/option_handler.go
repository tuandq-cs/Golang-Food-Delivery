package asyncjob

import "time"

type OptionHandler func(config *jobConfig)

func WithName(name string) OptionHandler {
	return func(config *jobConfig) {
		config.Name = name
	}
}

func WithRetryDurations(retries []time.Duration) OptionHandler {
	return func(config *jobConfig) {
		config.Retries = retries
	}
}
