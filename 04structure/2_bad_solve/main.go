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
	var r rect = rect{1, 2, 3, 4}
	fmt.Println("rect area = ", rectArea(r.x1, r.y1, r.x2, r.y2))
	//point
	var p point = point{0, 2}
	if pointInRect(r.x1, r.y1, r.x2, r.y2, p.x, p.y) {
		fmt.Println("point in rect")
	} else {
		fmt.Println("point outside rect")
	}
	structCreationExample()
}

func distance(x1, y1, x2, y2 float64) float64 {
	xd := x1 - x2
	yd := y1 - y2
	return math.Sqrt(xd*xd + yd*yd)
}

func rectArea(x1, y1, x2, y2 float64) float64 {
	width := distance(x1, y1, x2, y1)
	height := distance(x1, y1, x1, y2)
	return width * height
}

func pointInRect(x1, y1, x2, y2, px1, py1 float64) bool {
	maxX := math.Max(x1, x2)
	minX := math.Min(x1, x2)
	maxY := math.Max(y1, y2)
	minY := math.Min(y1, y2)
	if px1 > maxX || px1 < minX || py1 > maxY || py1 < minY {
		return false
	}
	return true
}
