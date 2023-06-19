package entity

type Book struct {
	Order []*Order
	Transaction []*Transaction
	OrderChan chan *Order
	OrderChanOut 
