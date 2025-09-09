package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//Создаю канал для записи
	dataChan := make(chan int)
	var wg sync.WaitGroup
	var a time.Duration
	fmt.Println("Введи время работы приложения в секундах:")
	fmt.Scanln(&a)
	timeOut := a * time.Second
	//Создаю отмену контекса по времени
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	//Горутина которая пишет в канал
	go func() {
		i := 0
		for {
			dataChan <- i
			i++
			time.Sleep(200 * time.Millisecond)
		}
	}()
	//Вейт группа для корректного завершения программы
	wg.Add(1)
	//Запускаю воркера
	go Worker(ctx, dataChan, &wg)
	wg.Wait()
}

func Worker(ctx context.Context, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		//отменяю контекст и завершаю работу воркера
		case <-ctx.Done():
			fmt.Println("Воркер завершил работу")
			return
		default:
			res := <-dataChan
			fmt.Println(res)
		}
	}
}
