package main

import (
	"fmt"
	"time"
)

type ControlMsg int

const (
	DoExit = iota
	ExitOK
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

func doubler(jobs <-chan Job, results chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			// only receive one message that it's ok to exit
			case DoExit:
				fmt.Println("Exit goroutine")
				control <- ExitOK
				return
			default:
				panic("unhandled control message")
			}
		case job := <-jobs:
			results <- Result{result: job.data * 2, job: job}
		}
	}
}

func main() {

	// jobs
	jobs := make(chan Job, 5000)
	// results
	results := make(chan Result, 50)
	// control
	control := make(chan ControlMsg)

	go doubler(jobs, results, control)

	for i := range 4999 {
		jobs <- Job{i}
	}

	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timed out")
			control <- DoExit
			finalMsg := <-control
			fmt.Println("Received final msg:", finalMsg, "program exit")
			return
		}
	}
}
