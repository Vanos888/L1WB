package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	//Отмена через условие
	wg.Add(1)
	stop := "stop"
	go func() {
		defer wg.Done()
		defer fmt.Println("Горутина 1 завершилась")
		if stop == "stop" {
			return
		}
	}()

	//Отмена через канал
	wg.Add(1)
	chanStop := make(chan string)
	go func() {
		defer wg.Done()
		fmt.Println("Горутина 2 завершилась")
		<-chanStop
	}()
	chanStop <- "stop"

	//Отмена через контекст
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("Горутина 3 завершилась")
		default:
		}
	}()
	cancel()

	//Отмена через закрытие канала
	wg.Add(1)
	closingChannel := make(chan int)
	go func() {
		defer wg.Done()
		for {
			if _, ok := <-closingChannel; !ok {
				fmt.Println("Горутина 4 завершилась")
				return
			}
		}
	}()
	close(closingChannel)
	//Отмена через runtime.Goexit
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Горутина 5 завершилась")
		runtime.Goexit()
	}()

	wg.Wait()
}
