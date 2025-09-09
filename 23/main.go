package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	target := 6

	for i := range arr {
		if arr[i] == target {
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	fmt.Println(arr)
}
