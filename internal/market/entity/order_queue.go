package entity

type OrderQueue struct {
	Orders []*Order
}

// Less reports whether the element with
func (oq *OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

// Swap swaps the elements with indexes i and j.
func (oq *OrderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

// Len is the number of elements in the collection.
func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

// Push pushes the element x onto the heap.
func (oq *OrderQueue) Push(x any) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

// Pop removes and returns the minimum element (according to Less) from the heap.
func (oq *OrderQueue) Pop() any {
	old := oq.Orders
	n := len(old)
	x := old[n-1]
	oq.Orders = old[0 : n-1]
	return x
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{
		Orders: []*Order{},
	}
}
