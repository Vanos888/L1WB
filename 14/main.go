package main

import "fmt"

func main() {
	TypeIs(10)
	TypeIs("10")
	TypeIs(true)
	TypeIs(make(chan int))
}

func TypeIs(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("chan")
	}
}
