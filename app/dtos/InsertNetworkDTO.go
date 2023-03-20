package dtos

type InsertNetworkDTO struct {
	Name        string `json:"name"`
	ShortName   string `json:"shortName"`
	NetworkType string `json:"networkType"`
	ChainType   uint   `json:"chainType"`
	Api         string `json:"api"`
	ScanUrl     string `json:"scanUrl"`
}
