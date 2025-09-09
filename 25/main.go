package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	sleep(2 * time.Second)
	fmt.Println("Hello World")

}
func sleep(t time.Duration) {
	<-time.After(t)
}
