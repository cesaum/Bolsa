package entity

type Investors struct {
	ID string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInver

type InvestorAssetPosition struct{
	AssetID string
	Shares int
	}
	
