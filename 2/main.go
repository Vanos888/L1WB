package main

import (
	"fmt"
	"sync"
)

// Функция для подсчета квадратов переданных чисел
func Calculate(number int, wg *sync.WaitGroup) {
	res := number * number
	defer wg.Done()
	fmt.Println(number, "*", number, "=", res)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	//Прохожусь слайсу и запускаю горутину для каждого числа в слайсе
	for _, number := range numbers {
		wg.Add(1)

		go Calculate(number, &wg)

	}
	//Жду завершения всех горутин
	wg.Wait()

}
