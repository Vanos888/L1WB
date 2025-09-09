package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Проблема кода в том, что "justString" ссылался на исходную строку "v" из-за этого возникала проблема с утечкой памяти
//т.к мы сохраняем всю исходную строку "v"
//Ниже представлен исправленный код, где мы копируем только нужную часть исходной строки,
//а исходная строка "v" может быть убрана сборщиком мусора

var justString string

func main() {
	someFunc()
}

func someFunc() {
	v := createHugeString(1024)
	justString = string([]byte(v[:100]))
	fmt.Println(v)
	fmt.Println(justString)
}

func createHugeString(size int) string {
	rand.Seed(time.Now().UnixNano())

	res := make([]byte, size)

	for i := 0; i < size; i++ {
		res[i] = byte(rand.Intn(95) + 32)
	}

	return string(res)
}
