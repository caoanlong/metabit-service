package repositories

import (
	"gorm.io/gorm"
	"metabit-service/core/config"
	"metabit-service/core/models"
)

type SysUserRepository interface {
	FindAll() []models.SysUser
	FindList() []models.SysUser
	FindById(sysUserID uint64) models.SysUser
	Insert(sysUser models.SysUser) models.SysUser
	Update(sysUser models.SysUser) models.SysUser
	Del(sysUser models.SysUser)
}

type sysUserConnection struct {
	connection *gorm.DB
}

func NewSysUserRepository() SysUserRepository {
	return &sysUserConnection{
		connection: config.G_DB,
	}
}

func (db *sysUserConnection) FindAll() []models.SysUser {
	var sysUsers []models.SysUser
	db.connection.Find(&sysUsers)
	return sysUsers
}

func (db *sysUserConnection) FindList() []models.SysUser {
	var sysUsers []models.SysUser
	db.connection.Find(&sysUsers)
	return sysUsers
}

func (db *sysUserConnection) FindById(sysUserID uint64) models.SysUser {
	var sysUser models.SysUser
	db.connection.Find(&sysUser, sysUserID)
	return sysUser
}

func (db *sysUserConnection) Insert(sysUser models.SysUser) models.SysUser {
	db.connection.Save(&sysUser)
	return sysUser
}

func (db *sysUserConnection) Update(sysUser models.SysUser) models.SysUser {
	db.connection.Save(&sysUser)
	return sysUser
}

func (db *sysUserConnection) Del(sysUser models.SysUser) {
	db.connection.Delete(sysUser)
}
