package dtos

type InsertTokenDTO struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Logo     string `json:"logo"`
	Decimals uint8  `json:"decimals"`
	Address  string `json:"address"`
	Type     string `json:"string"`
	Category string `json:"category"`
}
