package models

type Token struct {
	ID       uint64 `json:"id" gorm:"primary_key:auto_increment"`
	Name     string `json:"name" gorm:"type:varchar(32)"`
	Symbol   string `json:"symbol" gorm:"type:varchar(32)"`
	Logo     string `json:"logo" gorm:"type:varchar(256)"`
	Decimals uint8  `json:"decimals" gorm:"type:int(8)"`
	Address  string `json:"address" gorm:"type:varchar(128)"`
	Balance  string `json:"balance" gorm:"type:varchar(32)"`
	Type     string `json:"string" gorm:"type:varchar(32)"`
	Category string `json:"category" gorm:"type:varchar(32)"`
}
