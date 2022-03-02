package asyncjob

import (
	"Golang_Edu/common"
	"context"
	"log"
)

type JobGroup struct {
	jobs         []Job
	isConcurrent bool
}

func NewGroup(isConcurrent bool, jobs ...Job) *JobGroup {
	return &JobGroup{
		jobs:         jobs,
		isConcurrent: isConcurrent,
	}
}

func (group *JobGroup) Run(context context.Context) error {
	jobs := group.jobs
	errChan := make(chan error, len(jobs))
	for i := range jobs {
		job := jobs[i]
		if group.isConcurrent {
			go func(aj Job) {
				defer common.Recovery()
				errChan <- group.runJob(context, aj)
			}(job)
			continue
		}
		err := group.runJob(context, job)
		if err != nil {
			return err
		}
	}
	for i := 0; i < len(jobs); i++ {
		if v := <-errChan; v != nil {
			return v
		}
	}
	return nil
}

func (group *JobGroup) runJob(context context.Context, job Job) error {
	if err := job.Execute(context); err != nil {
		for {
			log.Println(err)
			if job.State() == StateRetryFailed {
				return err
			}
			if job.Retry(context) == nil {
				return nil
			}

		}
	}
	return nil
}
