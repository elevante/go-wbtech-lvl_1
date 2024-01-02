package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) Distance(q *Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)

	fmt.Println("Distance between points :", p1.Distance(p2))
}
