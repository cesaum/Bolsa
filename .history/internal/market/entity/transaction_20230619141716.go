package entity

type Transaction struct {
	ID string
	SellingOrder *Order
	BuyingOrder *Order
	Shares int
	Prince float64
	Total float64
	
}