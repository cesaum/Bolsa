package entity

import (
	"sync"
	"container/heap"
	)

type Book struct {
	Order        []*Order
	Transaction  []*Transaction
	OrdersChan    chan *Order
	OrderChanOut chan *Order
	Wg           *sync.WaitGroup
}

func NewBook(orderChan chan *Order, orderChanOut chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		Order:        []*Order{},
		Transaction:  []*Transaction{},
		OrderChan:    orderChan,
		OrderChanOut: orderChan,
		Wg:           wg,
	}
}

func (b *Book) Trade() {
	buyOrders := NewOrderQueue()
	sellOrders := NewOrderQueue()

	heap.Init(buyOrders)
	heap.Init(sellOrders)

	for order := range b.OrdersChan {
		if order.OrderType == "BUY"{
			buyOrders.Push(or)
		}
}
