package main

import "image/color"

type Points struct {
	X, Y float64
}

type ColorPoint struct {
	Points
	Color color.RGBA
}

func (p Points) Add(q Points) Points {
	return Points{p.X + q.X, p.Y + q.Y}
}

func (p Points) Sub(q Points) Points {
	return Points{p.X - q.X, p.Y - q.Y}
}

type Paths []Points

func (path Paths) translateBy(offset Points, add bool) {
	var op func(p, q Points) Points
	if add {
		op = Points.Add
	} else {
		op = Points.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}

}

func main() {

}
