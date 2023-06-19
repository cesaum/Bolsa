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

func NewTransaction(SellingOrder *Order, B