package entity

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price       float64
	Total        float64
	DataTime     time.Time
}

func NewTransaction(sellingOrder *Order, buyingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price
	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Shares:       shares,
		Price:       price,
		Total:        total,
		DataTime:     time.Now(),
	}
}


func (t *Transaction) CalculateTotalPrice(shares int, price float64) {
	t.Total = float64(t.Shares) * t.Price
}

func (t *Transaction) ClosedBuyTransaction() {
	if transaction.BuyingOrder.Pending == 0 {
		transaction.BuyingOrder.Pending.Order.Status = "CLOSED"		
	}
}
