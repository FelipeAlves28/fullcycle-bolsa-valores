package entity

type Asset struct {
	ID           string
	Name         string
	MarketVolume int
}

func NewAsset(id string, name string) *Asset {
	return &Asset{
		ID:   id,
		Name: name,
	}
}
