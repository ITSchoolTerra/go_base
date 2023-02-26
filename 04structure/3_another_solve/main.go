package main

import (
	"fmt"
	"math"
)

type rect struct {
	x1, y1, x2, y2 float64
}
type point struct {
	x, y float64
}

func main() {
	//rect
	var r rect = rect{0, 2, 8, 4}
	fmt.Println("rect area = ", rectArea(r))
	//point
	var p point = point{0, 2}
	if pointInRect(r, p) {
		fmt.Println("point in rect")
	} else {
		fmt.Println("point outside rect")
	}

}

func distance(x1, y1, x2, y2 float64) float64 {
	xd := x1 - x2
	yd := y1 - y2
	return math.Sqrt(xd*xd + yd*yd)
}

func rectArea(r rect) float64 {
	width := distance(r.x1, r.y1, r.x2, r.y1)
	height := distance(r.x1, r.y1, r.x1, r.y2)
	return width * height
}

func pointInRect(r rect, p point) bool {
	maxX := math.Max(r.x1, r.x2)
	minX := math.Min(r.x1, r.x2)
	maxY := math.Max(r.y1, r.y2)
	minY := math.Min(r.y1, r.y2)
	if p.x > maxX || p.x < minX || p.x > maxY || p.y < minY {
		return false
	}
	return true
}
