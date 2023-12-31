package entity

type Investor struct {
	ID string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID: id,
		AssetPosition: []*InvestorAssetPosition{},//arr
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition){
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

type InvestorAssetPosition struct{
	AssetID string
	Shares int
	}
	
