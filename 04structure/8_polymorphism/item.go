package main

import "fmt"

type usableItemInterface interface {
	ToString() string
	Use()
}
type item struct {
	price float64
	name  string
}

func (i *item) ToString() string {
	return i.name
}

type weapon struct {
	item
	damage float64
}

func (w *weapon) Use() {
	fmt.Printf("Нанесено %v урона\n", w.damage)

}

type potion struct {
	item
	power float64
}

func (p *potion) Use() {
	fmt.Printf("Восстановленно %v здоровья\n", p.power)

}
