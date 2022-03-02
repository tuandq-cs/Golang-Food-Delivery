package main

import (
	"Golang_Edu/component/asyncjob"
	"context"
	"log"
	"time"
)

func main() {
	job1 := asyncjob.NewJob(func(context context.Context) error {
		time.Sleep(time.Second)
		log.Println("Job 1 finished")
		return nil
		//return errors.New("job 1 crash")
	}, asyncjob.WithName("Job1"))
	job2 := asyncjob.NewJob(func(context context.Context) error {
		time.Sleep(time.Second * 2)
		log.Println("Job 2 finished")
		return nil
	}, asyncjob.WithName("Job2"))
	group := asyncjob.NewGroup(true, job1, job2)
	if err := group.Run(context.Background()); err != nil {
		log.Println(err)
	}
}
