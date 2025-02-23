package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(1 * time.Second) // Simulate work
		results <- job * 2          // Send result
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close jobs channel to indicate no more jobs

	// Collect results
	for r := 1; r <= numJobs; r++ {
		fmt.Println("Result:", <-results)
	}
}
