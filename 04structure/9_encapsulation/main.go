package main

import (
	"fmt"
	"structure/9_encapsulation/item"
)

func main() {
	Itm := item.Item{10, "Item"}
	fmt.Println(Itm.ToString())
	fmt.Println(Itm.Name)
	itm := item.NewItem(15, "item")
	fmt.Println(itm.ToString())
	fmt.Println(itm.Price)
	//fmt.Println(itm.name)
	//itm = item.item{10, "Item"}
	//itm = item.item{}
	PItm := item.PrivateItem{}
	PItm = item.PrivateItem{10, "PrivateItem"}
	fmt.Println(PItm.toString())
}
