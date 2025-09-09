package main

import (
	"fmt"
	"sync"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	number := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range number {
			chan1 <- v
		}
		close(chan1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range chan1 {
			chan2 <- i * 2
		}
		close(chan2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range chan2 {
			fmt.Println(i)
		}
	}()
	wg.Wait()

}
