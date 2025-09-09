package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu    sync.Mutex
}

func main() {
	var c Counter
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go c.inc(&wg)
	}
	wg.Wait()

	fmt.Println(c.upCount())
}

func (c *Counter) inc(wg *sync.WaitGroup) {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.count++
	wg.Done()
}

func (c *Counter) upCount() int {
	defer c.mu.Unlock()
	c.mu.Lock()
	return c.count
}
