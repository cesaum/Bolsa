package entity

type Investors struct {
	ID string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvertor(id string) *Invertor {
	return
}

type InvestorAssetPosition struct{
	AssetID string
	Shares int
	}
	
