package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go showGoRoutine(i, wg)
	}
	wg.Wait()
}

func showGoRoutine(id int, wg *sync.WaitGroup) {
	delay := rand.Intn(500)
	fmt.Printf("Go routine #%d\n", id)

	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}
