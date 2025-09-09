package main

import (
	"fmt"
)

func main() {
	var num int64 = 5
	var bitPos uint = 1

	res := Set(num, bitPos, 0)
	fmt.Printf("%d (%04b) после установки %d-го бита в 0: %d (%04b)\n",
		num, num, bitPos, res, res)

	res = Set(res, bitPos, 1)
	fmt.Printf("%d (%04b) после установки %d-го бита в 1: %d (%04b)\n",
		res&^(1<<bitPos), res&^(1<<bitPos), bitPos, res, res)
}

func Set(value int64, i uint, bitValue int) int64 {
	if bitValue == 1 {
		return value | (1 << i)
	} else {
		return value &^ (1 << i)
	}
}
