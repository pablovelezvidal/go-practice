package main

import (
	"fmt"
)

func main() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)

	Print(doubles)
}

func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func Double(in <-chan int, out chan<- int) { // canal escritura, flechita a la derecha de chan, lectura va a la izquierda
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

func Print(out chan int) {
	for value := range out {
		fmt.Println(value)
	}
}
