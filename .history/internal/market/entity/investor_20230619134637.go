package entity

type Investors struct {
	ID string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID: id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

func

type InvestorAssetPosition struct{
	AssetID string
	Shares int
	}
	
