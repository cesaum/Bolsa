package entity

import (
	"container/heap"
	"sync"
)

type Book struct {
	Order        []*Order
	Transaction  []*Transaction
	OrdersChan   chan *Order
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
		if order.OrderType == "BUY" {
			buyOrders.Push(order)
			if sellOrders.Len() > 0 && sellOrders.Len() < int(sellOrders.Orders[0].Price) <= order.Price {
				sellOrder := sellOrders.Pop().(*Order)
				if sellOrder.PendingShares > 0 {
					transaction 
					} 
			}
		}
	}
}
