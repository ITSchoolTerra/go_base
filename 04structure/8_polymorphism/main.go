package main

import "fmt"

func main() {
	sword := weapon{item{30, "iron sword"}, 5}
	healthPotion := potion{item{30, "iron sword"},
		10.5}
	player := character{}
	fmt.Println(player.ToString())
	player.Use()
	player.hands = &sword
	fmt.Println(player.ToString())
	player.Use()
	player.hands = &healthPotion
	fmt.Println(player.ToString())
	player.Use()
}
