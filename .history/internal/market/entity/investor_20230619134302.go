package entity

type Investors struct {
	ID string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvertor(id string) *Invertor {
	return &Invertor{
		ID:
	}

	
}

type InvestorAssetPosition struct{
	AssetID string
	Shares int
	}
	
