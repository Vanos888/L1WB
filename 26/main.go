package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abCsdhjrY"
	fmt.Println(Unique(str))
}

func Unique(str string) bool {
	lower := strings.ToLower(str)
	r := []rune(lower)

	for i := range r {
		for j := 1; j < len(r); j++ {
			if r[i] == r[j] && j != i {
				return false
			}
		}
	}
	return true
}
