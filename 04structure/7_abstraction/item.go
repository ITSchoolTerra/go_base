package main

type itemInterface interface {
	GetPrice() float64
	ToString() string
}
type item struct {
	price float64
	name  string
}

func (i *item) GetPrice() float64 {
	return i.price
}

type weapon struct {
	item
	damage float64
	name   string
}

func (i *item) ToString() string {
	return i.name + ", "
}
