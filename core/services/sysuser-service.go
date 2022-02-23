package services

import (
	"metabit-service/core/models"
	"metabit-service/core/repositories"
)

type SysUserService interface {
	FindAll() []models.SysUser
	FindList() []models.SysUser
	FindById(sysUserID uint64) models.SysUser
	Insert(sysUser models.SysUser) models.SysUser
	Update(sysUser models.SysUser) models.SysUser
	Del(sysUserID uint64)
	Login(sysUser models.SysUser)
}

type sysUserService struct {
	sysUserRepository repositories.SysUserRepository
}

func NewSysUserService() SysUserService {
	return &sysUserService{
		sysUserRepository: repositories.NewSysUserRepository(),
	}
}

func (service *sysUserService) FindAll() []models.SysUser {
	return service.sysUserRepository.FindAll()
}

func (service *sysUserService) FindList() []models.SysUser {
	//TODO implement me
	panic("implement me")
}

func (service *sysUserService) FindById(sysUserID uint64) models.SysUser {
	return service.sysUserRepository.FindById(sysUserID)
}

func (service *sysUserService) Insert(sysUser models.SysUser) models.SysUser {
	return service.sysUserRepository.Insert(sysUser)
}

func (service *sysUserService) Update(sysUser models.SysUser) models.SysUser {
	//TODO implement me
	panic("implement me")
}

func (service *sysUserService) Del(sysUserID uint64) {
	//TODO implement me
	panic("implement me")
}

func (service *sysUserService) Login(sysUser models.SysUser) {

	panic("implement me")
}
