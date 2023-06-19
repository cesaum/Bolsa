package entity

import (
	"time"
	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Prince       float64
	Total        float64
	DataTime     time.Time
}

func NewTransaction(sellingOrder *Order, buyingOrder *Order,Shares int, price float64) *Transaction {
	total := float64(shares) * price
	return &Transaction{
		ID:uuid.New().string(),
		SellingOrder: sellingOrder,	
		BuyingOrder: buyingOrder,	
		Shares: shares,
		Prince: prices,
		Total: total,
		DataTime: ,

	}
