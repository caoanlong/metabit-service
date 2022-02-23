package models

type SysUser struct {
	Base
	NickName string `json:"nickname" gorm:"comment:昵称"`
	Username string `json:"username" gorm:"comment:用户名"`
	Email    string `json:"email" gorm:"comment:邮箱"`
	Phone    string `json:"phone" gorm:"comment:手机"`
	Password string `json:"password" gorm:"comment:密码"`
	Salt     string `json:"salt" gorm:"comment:盐"`
}
