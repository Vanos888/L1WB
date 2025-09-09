package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	dataChan := make(chan int)
	quantityWorkers := 3
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		counter := 0

		for {
			dataChan <- counter
			counter++
			time.Sleep(200 * time.Millisecond)
		}
	}()

	for i := 1; i < quantityWorkers+1; i++ {
		wg.Add(1)
		go Workers(ctx, dataChan, i, &wg)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		<-sigChan
		cancel()
	}()

	wg.Wait()
	close(dataChan)
}

func Workers(ctx context.Context, dataChan chan int, numberWorker int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():

			fmt.Printf("Воркер %d завершил работу\n", numberWorker)
			return
		default:
			res := <-dataChan
			fmt.Printf("Воркер %v отработал, результат: %d \n", numberWorker, res)
		}
	}
}
