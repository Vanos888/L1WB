package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Создаю канал интов
	dataChan := make(chan int)
	var wg sync.WaitGroup
	//Создаю переменную в которую передаю количество воркеров
	var quantityWorkers int
	fmt.Println("Введите количество воркеров: ")
	fmt.Scanln(&quantityWorkers)

	//Пишу в канал в отдельной горутине
	go func() {
		counter := 0

		for {
			dataChan <- counter
			counter++
			time.Sleep(200 * time.Millisecond)
		}
	}()

	//Запускаю Воркеров
	for i := 1; i < quantityWorkers+1; i++ {
		wg.Add(1)
		go Workers(dataChan, i)
	}
	//Жду завершения всех горутин и закрываю канал
	wg.Wait()
	close(dataChan)
}

func Workers(dataChan chan int, numberWorker int) {
	for {
		//Читаю из канала и вывожу в STDOUT
		res := <-dataChan
		fmt.Printf("Воркер %v отработал, результат: %d \n", numberWorker, res)
	}
}
