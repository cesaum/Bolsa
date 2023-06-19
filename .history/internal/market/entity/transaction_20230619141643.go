package entity

type Transaction struct {
	ID string
	SellingOrder *Order
	By
}
