package main

import "fmt"

func main() {
	a := []int{1, 2, 7, 3, 4, -1}
	b := []int{2, 3, 5, 6, 4, -1, 7}
	var res []int

	for _, i := range a {
		for _, j := range b {
			if i == j {
				res = append(res, i)
			}
		}
	}
	fmt.Println(res)
}
