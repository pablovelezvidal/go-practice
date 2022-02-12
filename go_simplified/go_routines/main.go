package main

import (
	"fmt"
	"go_routines/data"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// executeShowGoRoutine()
	start := time.Now()
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readBook(i, wg, m)
	}
	wg.Wait()
	duration := time.Since(start).Milliseconds()

	fmt.Printf("Reading books duration: %d\n", duration)

}

func readBook(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
	data.FinishBook(id, m)

	delay := rand.Intn(800)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}

// func executeShowGoRoutine() {
// 	start := time.Now()
// 	wg := &sync.WaitGroup{}
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go showGoRoutine(i, wg)
// 	}
// 	wg.Wait()
// 	duration := time.Since(start).Milliseconds()
// 	fmt.Printf("Duration %d", duration)
// }

// func showGoRoutine(id int, wg *sync.WaitGroup) {
// 	delay := rand.Intn(500)
// 	fmt.Printf("Go routine #%d\n", id)

// 	time.Sleep(time.Millisecond * time.Duration(delay))
// 	wg.Done()
// }
