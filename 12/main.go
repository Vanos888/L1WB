package main

import "fmt"

func main() {
	example := []string{"cat", "cat", "dog", "cat", "tree", "asdasd", "cat", "dog"}

	var res []string

	for _, i := range example {
		unique := false

		for _, j := range res {
			if i == j {
				unique = true
				break
			}
		}
		if unique == false {
			res = append(res, i)
		}
	}
	fmt.Println(res)
}
