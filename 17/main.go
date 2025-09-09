package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 6} //8
	target := 3
	fmt.Println(BSearch(arr, target))
}

func BSearch(arr []int, target int) int {
	core := len(arr) / 2
	right := len(arr) - 1
	if target == arr[right] {
		return right
	}
	for {
		if target < arr[core] {
			right = core
			core = core / 2
			continue
		}
		if target > arr[core] {
			core = ((right - core) / 2) + core
			continue
		}
		if target == arr[core] {
			return core

		}
		break
	}
	return 0
}
