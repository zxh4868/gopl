package main

type point struct {
	x, y int
}

type circle struct {
	center point
	radius float64
}

type wheel struct {
	circle circle
	spots  point
}

func main() {
	var w wheel
	w.spots.x = 1
	w.spots.y = 2
	w.circle.center.x = 3
	w.circle.center.y = 4
	w.circle.radius = 5

}
