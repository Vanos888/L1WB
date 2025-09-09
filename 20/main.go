package main

import "fmt"

func main() {
	str := "я вчера поймал большую щуку"
	b := []byte(str)

	reverse(b, 0, len(b)-1)

	a := 0

	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b, a, i-1)
			a = i + 1
		}
	}
	fmt.Println(string(b))

}

func reverse(byte []byte, a int, b int) {
	for i, j := a, b; i < j; i, j = i+1, j-1 {
		byte[i], byte[j] = byte[j], byte[i]
	}
}
