package main

import (
	"fmt"
	"sync"
)

func main() {
	var score = []int{0}
	wg := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("First GoRoutine")
		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Second GoRoutine")
		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Third GoRoutine")
		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	// go func(wg *sync.WaitGroup, m *sync.RWMutex) {
	// 	fmt.Println("Read GoRoutine")
	// 	mutex.RLock()
	// 	fmt.Println(score)
	// 	mutex.RUnlock()
	// 	wg.Done()
	// }(wg, mutex)

	wg.Wait()
	fmt.Println(score)
}
