package main

import "fmt"

const (
	itemT   = 0
	weaponT = 1
)

type itemType int

func (it itemType) ToString() string {
	switch it {
	case 1:
		return "Оружие"
	default:
		return "Предмет"
	}
}

type item struct {
	price    float64
	name     string
	itemType itemType
}

func (i *item) SetPrice(p float64) {
	i.price = p
}

func (i *item) ToString() string {
	return fmt.Sprintf("%v %v price: %v",
		i.itemType.ToString(), i.name, i.price)
}

type weapon struct {
	item
	damage float64
	name   string
}

func (w *weapon) ToString() string {
	return fmt.Sprintf("%v %v(item name %v) "+
		"price: %v damage: %v",
		w.itemType.ToString(), w.name, w.item.name,
		w.price, w.damage)
}