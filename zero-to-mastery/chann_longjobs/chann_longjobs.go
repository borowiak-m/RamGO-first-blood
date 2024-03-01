package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job int

func longCalculation(job Job) int {
	duration := time.Duration(rand.Intn(10000)) * time.Millisecond
	time.Sleep(duration)
	fmt.Printf("Job %d complete in %v\n", job, duration)
	return int(job) * 30
}

func runJob(resultChan chan int, job Job) {
	resultChan <- longCalculation(job)
}

func makeJobs() []Job {
	jobs := make([]Job, 0, 100)
	for i := 0; i < 100; i++ {
		jobs = append(jobs, Job(rand.Intn(10000)))
	}
	return jobs
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// create a queue of jobs
	jobs := makeJobs()
	// create a channel for results
	resultChan := make(chan int, 10)
	// load results channel with processed jobs results per iteration
	for i := 0; i < len(jobs); i++ {
		// spawn separate processes to handle a job per exec
		go runJob(resultChan, jobs[i])
	}

	resultCount := 0
	sum := 0

	for {
		// receive processing results from the results channel
		result := <-resultChan
		// results are ints, sum them up
		sum += result
		// count how many results processed
		resultCount++
		// if as many as jobs created, end loop
		if resultCount == len(jobs) {
			break
		}
	}
	fmt.Println("sum is", sum)
}
