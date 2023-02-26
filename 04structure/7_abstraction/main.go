package main

import "fmt"

func main() {
	secretMap := item{20, "secret map"}
	sword := weapon{item{30, "iron sword"},
		5, "rusty sword"}
	player := character{gold: 400, items: []itemInterface{&sword}}
	fmt.Println(player.ToString())
	player.buyItem(secretMap)
	player.buyWeapon(sword)
	fmt.Println(player.ToString())
	player.buy2(sword)
	fmt.Println(player.ToString())
	player.buy3(&sword)
	fmt.Println(player.ToString())
}
