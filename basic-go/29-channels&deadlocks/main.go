package main

import (
	"fmt"
	"sync"
)

func main() {
	myChannel := make(chan int, 1)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChannel

		fmt.Println(val)
		fmt.Println(isChannelOpen)

		wg.Done()
	}(myChannel, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 1
		close(myChannel)

		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
