package models

import (
	"time"
)

type Base struct {
	ID        uint       `json:"id" gorm:"primarykey"` // 主键ID
	CreatedAt *time.Time `json:"createdAt"`            // 创建时间
	UpdatedAt *time.Time `json:"updatedAt"`            // 更新时间
	DeletedAt *time.Time `json:"deletedAt"`            // 删除时间
}
