package main

import (
	"fmt"
	"sync"
)

func mergeChans[T any](ch1, ch2 <-chan T) <-chan T {
	res := make(chan T)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range ch1 {
			res <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range ch2 {
			res <- v
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(res)
	}()
	return res
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 10; i < 20; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	for v := range mergeChans(ch1, ch2) {
		fmt.Println(v)
	}
}
