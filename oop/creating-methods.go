package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Start, End Point
}

type Path []Point

func (l Line) GetDistance() float64 {
	return math.Hypot(l.End.X-l.Start.X, l.End.Y-l.Start.Y)
}

func (p Path) GetDistance() (sum float64) {
	// `sum` here is just a local var
	for i := 1; i < len(p); i++ {
		sum += Line{p[i-1], p[i]}.GetDistance()
	}
	return sum
}

type N string

func (n N) GetName() {
	fmt.Println("Hello,", n)
}

func main() {
	// simple method
	var name N = "Ahmad"
	name.GetName() // Hello, Ahmad

	someLine := Line{Point{5, 9}, Point{30, 78}}
	somePath := Path{{1, 3}, {7, 4}, {8, 13}, {21, 30}}

	fmt.Println(someLine.GetDistance())
	fmt.Println(somePath.GetDistance())
}
