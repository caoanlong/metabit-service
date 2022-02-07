package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"metabit-service/api/services"
	"metabit-service/api/utils"
)

type TokenController interface {
	FindAll(ctx *gin.Context)
	FindList(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Del(ctx *gin.Context)
}
type tokenController struct {
	tokenService services.TokenService
}

func NewTokenController() TokenController {
	return &tokenController{
		tokenService: services.NewTokenService(),
	}
}

func (t tokenController) FindAll(ctx *gin.Context) {
	utils.Success(ctx)
}

func (t tokenController) FindList(ctx *gin.Context) {
	//chain := ctx.Query("chain")
	env := viper.GetString("env")
	utils.Success(ctx, env)
}

func (t tokenController) FindById(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t tokenController) Insert(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t tokenController) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t tokenController) Del(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
