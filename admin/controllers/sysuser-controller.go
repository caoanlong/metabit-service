package controllers

import (
	"github.com/gin-gonic/gin"
	"metabit-service/admin/dtos"
	"metabit-service/core/models"
	"metabit-service/core/services"
	"metabit-service/core/utils"
)

type SysUserController interface {
	FindAll(ctx *gin.Context)
	FindList(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Del(ctx *gin.Context)
	Login(ctx *gin.Context)
}
type sysUserController struct {
	sysUserService services.SysUserService
}

func NewSysUserController() SysUserController {
	return &sysUserController{
		sysUserService: services.NewSysUserService(),
	}
}

func (controller *sysUserController) FindAll(ctx *gin.Context) {
	list := controller.sysUserService.FindAll()
	utils.Success(ctx, list)
}

func (controller *sysUserController) FindList(ctx *gin.Context) {
	//chain := ctx.Query("chain")
	//env := config.G_CONFIG.Env
	list := controller.sysUserService.FindList()
	utils.Success(ctx, list)
}

func (controller *sysUserController) FindById(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *sysUserController) Insert(ctx *gin.Context) {
	utils.Success(ctx)
}

func (controller *sysUserController) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *sysUserController) Del(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *sysUserController) Login(ctx *gin.Context) {
	var dto dtos.LoginSysUserDTO
	err := ctx.ShouldBind(&dto)
	if err != nil {
		utils.Fail(ctx, 500, err.Error())
	} else {
		sysUser := models.SysUser{
			Username: dto.Username,
			Password: dto.Password,
		}
		controller.sysUserService.Insert(sysUser)
		utils.Success(ctx)
	}
}
