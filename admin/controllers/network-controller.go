package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"metabit-service/core/services"
	"metabit-service/core/utils"
)

type NetworkController interface {
	FindAll(ctx *gin.Context)
	FindList(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Del(ctx *gin.Context)
}
type networkController struct {
	networkService services.NetworkService
}

func NewNetworkController() NetworkController {
	return &networkController{
		networkService: services.NewNetworkService(),
	}
}

func (t networkController) FindAll(ctx *gin.Context) {
	utils.Success(ctx)
}

func (t networkController) FindList(ctx *gin.Context) {
	//chain := ctx.Query("chain")
	env := viper.GetString("env")
	utils.Success(ctx, env)
}

func (t networkController) FindById(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t networkController) Insert(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t networkController) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t networkController) Del(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
