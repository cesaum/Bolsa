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

func (i *Investor) UpdateAssetPosition(assetID string, qtdShares int){
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, )
}

func (i *Investor) GetAssetPosition((assetID string) *InvestorAssetPosition{
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
}

type InvestorAssetPosition struct{
	AssetID string
	Shares int
	}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition {
	return &investorAssetPosition{
		AssetID: assetID,
		Shares: shares,
	}
}
