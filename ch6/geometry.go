package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}
type Path []Point

func (p Path) Distance() float64 {
	res := 0.0
	for i := range p {
		if i > 0 {
			res += p[i].Distance(p[i-1])
		}
	}
	return res
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {

	p := Point{3, 4}
	q := Point{5, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := Path{
		{1, 2},
		{4, 5},
		{3, 7},
		{1, 2},
	}

	fmt.Println("三角形的周长为：", perim.Distance())

}
