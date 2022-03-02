package asyncjob

import (
	"context"
	"log"
	"time"
)

type Job interface {
	Execute(context context.Context) error
	Retry(context context.Context) error
	State() JobState
	SetRetryDurations([]time.Duration)
}

type JobHandler func(context context.Context) error

type job struct {
	config     *jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler JobHandler, options ...OptionHandler) *job {
	newJob := job{
		config: &jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTimes,
		},
		handler:    handler,
		state:      StateInit,
		retryIndex: -1,
		stopChan:   make(chan bool),
	}
	for i := range options {
		options[i](newJob.config)
	}
	return &newJob
}

func (j *job) Execute(context context.Context) error {
	log.Printf("execute %s\n\n", j.config.Name)
	j.state = StateProcessing
	if err := j.handler(context); err != nil {
		j.state = StateFailed
		return err
	}
	j.state = StateCompleted
	return nil
}

func (j *job) Retry(context context.Context) error {
	j.retryIndex += 1
	time.Sleep(j.config.Retries[j.retryIndex])
	err := j.Execute(context)
	if err == nil {
		j.state = StateCompleted
	}
	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = StateRetryFailed
		return err
	}
	j.state = StateFailed
	return err
}

func (j *job) State() JobState {
	return j.state
}

func (j *job) SetRetryDurations(durations []time.Duration) {
	if len(durations) == 0 {
		return
	}
	j.config.Retries = durations
}
