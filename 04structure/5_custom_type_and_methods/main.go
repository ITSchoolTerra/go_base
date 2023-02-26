package main

import "fmt"

func main() {
	secretMap := item{20, "secret map", itemT}
	fmt.Println(secretMap)
	fmt.Println(secretMap.ToString())
	secretMap.SetPrice(5)
	fmt.Println(secretMap.ToString())
	sword := item{30, "iron sword", weaponT}
	fmt.Println(sword.ToString())
	sword.SetPricePointer(5)
	fmt.Println(sword.ToString())
}
