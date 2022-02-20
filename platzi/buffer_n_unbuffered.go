package main

import "fmt"

func main() {
	// c := make(chan int) // Unbuffered - expects a fn or routine to receive the message
	c := make(chan int, 2) // Buffered

	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)
}
