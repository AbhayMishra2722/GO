package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels")

	mych := make(chan int, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	//RECIEVE ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChannelOpen := <-mych
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		//fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)
	//SEND ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		//mych <- 5
		mych <- 0
		close(mych)
		wg.Done()
	}(mych, wg)
	wg.Wait()

	//mych <- 5
	//fmt.Println(<-mych)

}
