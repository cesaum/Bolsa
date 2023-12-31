package entity

type Investor struct {
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

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition)

type InvestorAssetPosition struct{
	AssetID string
	Shares int
	}
	
