package entity

type Investors struct {
	ID string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvertor

type InvestorAssetPosition struct{
	AssetID string
	Shares int
	}
	
