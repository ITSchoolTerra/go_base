package main

import "fmt"

type character struct {
	hands usableItemInterface
}

func (c *character) Use() {
	if c.hands == nil {
		fmt.Println("Нечего использовать")
		return
	}
	c.hands.Use()
}
func (c *character) ToString() string {
	if c.hands == nil {
		return "empty hands"
	}
	return fmt.Sprintf("in hands %v", c.hands.ToString())
}
