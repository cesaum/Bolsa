package entity

import "sync"
type Book struct {
	Order        []*Order
	Transaction  []*Transaction
	OrderChan    chan *Order
	OrderChanOut chan *Order
	Wg           *sync.WaitGroup
}

func NewBook(orderChan chan *Order, orderChanOut chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		Order: []*Order{},
		Transaction: []*Transaction{},
		OrderChan: orderChan,
		OrderChanOut: orderChan,
		Wg: wg,
	}
}

func (b *Book) Trade() {
	buyOrders :== NewOrderQu