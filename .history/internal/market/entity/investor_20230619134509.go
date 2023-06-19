package entity

type Investors struct {
	ID string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) *Invertor {
	return &Invertor{
		ID: id,
		AssetPosition: []*investorAssetPosition{},
	}
}

type InvestorAssetPosition struct{
	AssetID string
	Shares int
	}
	
