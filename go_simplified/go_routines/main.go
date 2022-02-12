package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go showGoRoutine(i)
	}
	time.Sleep(3 * time.Second)
}

func showGoRoutine(id int) {
	delay := rand.Intn(500)
	fmt.Printf("Go routine #%d\n", id)

	time.Sleep(time.Millisecond * time.Duration(delay))
}
