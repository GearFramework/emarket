package entities

type CartItem struct {
	ID       uint
	Quantity uint
}

type CartItems map[uint]CartItem

func (ci CartItems) Put(item CartItem) {
	if _, ok := ci[item.ID]; ok {
		i := ci[item.ID]
		i.Quantity++
		ci[item.ID] = i
		return
	}
	ci[item.ID] = item
}
