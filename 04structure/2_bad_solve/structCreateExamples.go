package main

import "fmt"

func structCreationExample() {
	var r = rect{0, 2, 4, 7}
	fmt.Println("r", r)
	r2 := rect{}
	fmt.Println("r2", r2)
	r2.x1 = -3
	r2.y1 = -1
	fmt.Println("r2", r2)
	r3 := new(rect)
	fmt.Println("r3", r3)
	r4 := newRect(0, 0, 3, 6)
	fmt.Println("r4", r4)
}

func newRect(x1in, y1in, x2in, y2in float64) rect {
	r5 := rect{
		x1: x1in,
		//y1: y1in,
		x2: x2in,
		y2: y2in,
	}
	return r5
}
