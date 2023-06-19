package entity

import "time"

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Prince       float64
	Total        float64
	DataTime     time.Time
}

func NewTransaction(sellingOrder *Order, buyingOrder *Order,Shares int, price float64, dataTime time.Time) *Transaction {