package item

type item struct {
	Price float64
	name  string
}

func NewItem(p float64, n string) item {
	return item{Price: p, name: n}
}

func (i *item) ToString() string {
	return i.name
}

type Item struct {
	Price float64
	Name  string
}

func (i *Item) ToString() string {
	return i.Name
}

type PrivateItem struct {
	price float64
	name  string
}

func (i *PrivateItem) toString() string {
	return i.name
}
