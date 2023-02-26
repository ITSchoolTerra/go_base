package main

import "fmt"

type character struct {
	gold  float64
	items []itemInterface
}

func (c *character) buyItem(i item) {
	if i.price <= c.gold {
		c.gold -= i.price
		c.items = append(c.items, &i)
	}
}
func (c *character) buyWeapon(i weapon) {
	if i.price <= c.gold {
		c.gold -= i.price
		c.items = append(c.items, &i)
	}
}
func (c *character) buy2(i interface{}) {
	switch i.(type) {
	case item:
		if i.(item).price <= c.gold {
			c.gold -= i.(item).price
			itm := i.(item)
			c.items = append(c.items, &itm)
		}
	case weapon:
		if i.(weapon).price <= c.gold {
			c.gold -= i.(weapon).price
			itm := i.(weapon)
			c.items = append(c.items, &itm)
		}
	}

}
func (c *character) buy3(i itemInterface) {
	if i.GetPrice() <= c.gold {
		c.gold -= i.GetPrice()
		c.items = append(c.items, i)
	}
}
func (c *character) ToString() string {
	res := fmt.Sprintf("gold: %v, ", c.gold)
	for _, i := range c.items {
		res += i.ToString()
	}
	return res
}
