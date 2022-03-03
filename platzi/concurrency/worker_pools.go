package main

import (
	"fmt"
)

func main() {
	tasks := []int{1, 2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Workers started with id %d \n", id)
		fib := Fibo(job)
		fmt.Printf("Worker finished, job: %d and fibo result %d \n", job, fib)
		results <- fib
	}
}

func Fibo(n int) int {
	if n <= 1 {
		return n
	}

	return Fibo(n-1) + Fibo(n-2)
}
