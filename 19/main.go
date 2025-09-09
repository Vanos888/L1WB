package main

import "fmt"

func main() {
	str := "ялюблюрыбалку"
	runes := []rune(str)
	var res []rune
	for i := len(runes) - 1; i >= 0; i-- {
		res = append(res, runes[i])
	}
	fmt.Println(string(res))
}
