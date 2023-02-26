package main

import (
	"fmt"
	"math"
)

func main() {
	//rect
	var x1, y1, x2, y2 float64 = 0, 0, 3, 4
	fmt.Println("rect area = ", rectArea(x1, y1, x2, y2))
	//point
	var px1, py1 float64 = 0, 2
	if pointInRect(x1, y1, x2, y2, px1, py1) {
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
