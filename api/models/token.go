package models

type Token struct {
	Base
	Name      string `json:"name" gorm:"comment:名称"`
	Symbol    string `json:"symbol" gorm:"comment:符号名"`
	Logo      string `json:"logo" gorm:"comment:LOGO"`
	Decimals  uint8  `json:"decimals" gorm:"comment:进制"`
	Address   string `json:"address" gorm:"comment:合约地址"`
	Balance   string `json:"balance" gorm:"comment:余额，默认为0"`
	ChainType string `json:"chainType" gorm:"comment:链类型"`
	Network   string `json:"network" gorm:"comment:网络"`
}
