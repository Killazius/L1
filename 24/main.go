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
func (p *Point) Distance(other *Point) float64 {
	xDiff := p.x - other.x
	yDiff := p.y - other.y
	return math.Sqrt(xDiff*xDiff + yDiff*yDiff)
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)
	fmt.Println(p1.Distance(p2))
}
