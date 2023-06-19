package entity

type Investor struct {
	ID string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID: id,
		AssetPosition: []*InvestorAssetPosition{},//array dinamico
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition){
	i.AssetPosition = append(i.AssetPosition, assetPosition) //adicionando o investor no array(slice)
}

func (i *Investor) UpdateAssetPosition(AssetID string, qtdShares int)
type InvestorAssetPosition struct{
	AssetID string
	Shares int
	}
	
