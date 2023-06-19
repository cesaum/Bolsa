package entity

type Order struct {
	ID string
	Investor *Investor
	Asset *Asset
	Shares int
	PendingShares int
	Price float64
}