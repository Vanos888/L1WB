package main

import "fmt"

func main() {
	human := Human{
		name: "Ivan",
		age:  24,
	}
	ac := Action{
		Human{
			name: "Vladimir",
			age:  25,
		}}
	human.GetName()
	ac.GetName()
}

type Human struct {
	name string
	age  int
}

func (h *Human) GetName() {
	fmt.Println("Привет, меня зовут", h.name)
}

type Action struct {
	Human
}
