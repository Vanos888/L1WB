package main

import "fmt"

func main() {
	arr := []int{5, 3, 9, 15, 2, 1, 26, 19, 35}

	sort := quickSort(arr)
	fmt.Println(sort)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	var res []int

	core := arr[len(arr)/2]

	left := []int{}
	right := []int{}
	middle := []int{}

	for _, num := range arr {
		if num < core {
			left = append(left, num)
			continue
		}
		if num > core {
			right = append(right, num)
			continue
		}
		middle = append(middle, num)

	}

	res = append(append(quickSort(left), middle...), quickSort(right)...)
	return res

}
