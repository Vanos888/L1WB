package main

import "fmt"

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 8, -5, 19.0, 15.5, 24.5, -21.0, 32.5}

	res := make(map[int][]float64)

	for _, v := range temp {

		group := 10 * int(v/10)
		res[group] = append(res[group], v)
	}
	fmt.Println(res)
}
