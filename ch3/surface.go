package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 300
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func corner(i, j int) (float64, float64) {
	// 求出网格单元的顶点坐标
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//计算曲面高度
	z := f(x, y)

	// 将(x,y,z)等角投射到二维SVG绘图平面上， 坐标是(sx,xy)
	sx := width/2 + (x-y)*cos30*xyscale
	xy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, xy

}
func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' >"+
		"style='stroke: grey; fill; white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygen point = '%g, %g, %g, %g, %g,%g,%g,%g'>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Printf("</svg>\n")
	
}
