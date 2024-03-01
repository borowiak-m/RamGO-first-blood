package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job int

func longCalculation(job Job) int {
	duration := time.Duration(rand.Intn(10000)) * time.Millisecond
	time.Sleep(duration)
	fmt.Printf("Job %d complete in %v\n", job, duration)
	return int(job) * 30
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

	resultCount := 0
	sum := 0

	// create a channel for results
	resultChan := make(chan int, 10)
	// define a wait group to monitor acive goroutine count
	var wg sync.WaitGroup
	// load results channel with processed jobs results per iteration
	for i := 0; i < len(jobs); i++ {
		wg.Add(1) // adding one goroutine
		// spawn separate processes to handle a job per exec
		go func(resultChan chan int, job Job) {
			defer wg.Done() // let waiting group know when finished
			resultChan <- longCalculation(job)
		}(resultChan, jobs[i])
	}

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
	wg.Wait() // not really needed here since we are looping until all results are read
	fmt.Println("sum is", sum)
}
