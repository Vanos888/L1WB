package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	data := map[int]int{}
	var mu sync.Mutex

	go func() {
		a := 1
		for {
			mu.Lock()
			data[a]++
			fmt.Println(data)
			time.Sleep(200 * time.Millisecond)
			mu.Unlock()
		}

	}()
	go func() {
		a := 2
		for {
			mu.Lock()
			data[a]++
			fmt.Println(data)
			time.Sleep(200 * time.Millisecond)
			mu.Unlock()
		}

	}()

	time.Sleep(5 * time.Second)
}
