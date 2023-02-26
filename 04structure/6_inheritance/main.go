package main

import "fmt"

func main() {
	secretMap := item{20, "secret map", 0}
	fmt.Println(secretMap.ToString())
	sword := weapon{item{30, "iron sword",
		1}, 5, "rusty sword"}
	fmt.Println(sword.ToString())
	fmt.Println(sword.item.ToString())
}
