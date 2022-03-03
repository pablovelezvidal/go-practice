package main

import (
	"fmt"
)

func main() {
	// controled defined number of tasks in this case is just integer numbers
	tasks := []int{1, 2, 3, 4, 5, 7, 10, 12, 40}
	// number of workers (processes that will be doing the tasks)
	nWorkers := 3
	// channel that will contain the information
	jobs := make(chan int, len(tasks))
	// channel that will store the output result of the process
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		// create the process that will run the tasks
		go Worker(i, jobs, results)
	}

	// fill the jobs process with the values
	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	// free the values from the results channel
	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	// iterate over the jobs channel
	for job := range jobs {
		fmt.Printf("Workers started with id %d \n", id)
		// do the work based on the values in the jobs channel
		fib := Fibo(job)
		fmt.Printf("Worker finished, job: %d and fibo result %d \n", job, fib)
		// send the reulting value to the results channel
		results <- fib
	}
}

func Fibo(n int) int {
	if n <= 1 {
		return n
	}

	return Fibo(n-1) + Fibo(n-2)
}
