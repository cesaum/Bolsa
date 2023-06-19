package entity

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	OrderType     string
	Status        string
	Transaction   []*Transaction
}

func NewOrder(orderID string, investor *Investor, asset *Asset, shares int, price float64, orderType string){
	return &Order{
		ID: orderID,
		Investor: investor,
		Asset: aasset,
		Shares: shares,
		PendingShares: shares,
		Price: price,
		Status: "OPEN",
		Transaction: ,
}