package entity

type Investors struct {
	ID string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Invertor{
		ID: id,
		AssetPosition: []*investorAssetPosition{},
	}
}

type InvestorAssetPosition struct{
	AssetID string
	Shares int
	}
	
