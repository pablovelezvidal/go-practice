package main

import (
	"fmt"
)

func main() {
	generator := make(chan int)
	doubles := make(chan int)

	// new process adding values to the generator channel
	go Generator(generator)
	// new process reading the generator values, doubling them and writing to the doubles channel
	go Double(generator, doubles)

	Print(doubles)
}

func Generator(c chan<- int) {
	// fill the generator channel with numbers
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func Double(in <-chan int, out chan<- int) { // canal escritura, flechita a la derecha de chan, lectura va a la izquierda
	// read in (generator channel) loop through it, double it and write it to the out channel (doubles)
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

func Print(out chan int) {
	// read all values in the output channel (doubles) and print it
	for value := range out {
		fmt.Println(value)
	}
}
