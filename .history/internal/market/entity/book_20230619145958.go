package entity

import "sync"
type Book struct {
	Order        []*Order
	Transaction  []*Transaction
	OrderChan    chan *Order
	OrderChanOut chan *Order
	Wg           *sync.WaitGroup
}

func NewBook
