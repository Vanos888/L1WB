package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 28)
	b := big.NewInt(1 << 25)

	fmt.Println(a.String(), "*", b, "=", a.Mul(a, b))
	fmt.Println(a.String(), "/", b, "=", a.Div(a, b))
	fmt.Println(a.String(), "+", b, "=", a.Add(a, b))
	fmt.Println(a.String(), "-", b, "=", a.Sub(a, b))
}
