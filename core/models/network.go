package models

type Network struct {
	Base
	Name        string `json:"name" gorm:"comment:名称，Ethereum Main Network"`
	ShortName   string `json:"shortName" gorm:"comment:短名称，Ethereum"`
	NetworkType string `json:"networkType" gorm:"comment:类型，mainnet"`
	ChainType   uint   `json:"chainType"`
	Api         string `json:"api" gorm:"comment:api地址"`
	ScanUrl     string `json:"scanUrl" gorm:"comment:scan地址"`
}
