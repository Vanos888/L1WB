package main

import (
	"fmt"
	"math"
)

type Point struct {
	a float64
	b float64
}

func NewPoint(a float64, b float64) *Point {
	return &Point{
		a: a,
		b: b,
	}
}

func main() {
	firstPoint := NewPoint(1, 2)
	secondPoint := NewPoint(6, 7)

	distance := firstPoint.Distance(secondPoint)
	fmt.Println(distance)

}

func (p *Point) Distance(q *Point) float64 {
	return math.Sqrt(math.Pow(p.a-q.a, 2) + math.Pow(p.b-q.b, 2))
}
