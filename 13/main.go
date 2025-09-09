package main

import "fmt"

func main() {
	a := 10
	b := 715
	b = a + b
	a = b - a
	b = b - a

	fmt.Println(a, b)
}
