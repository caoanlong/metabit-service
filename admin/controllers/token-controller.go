package controllers

import (
	"github.com/gin-gonic/gin"
	"metabit-service/admin/dtos"
	"metabit-service/core/models"
	"metabit-service/core/services"
	"metabit-service/core/utils"
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

func (controller *tokenController) FindAll(ctx *gin.Context) {
	list := controller.tokenService.FindAll()
	utils.Success(ctx, list)
}

func (controller *tokenController) FindList(ctx *gin.Context) {
	//chain := ctx.Query("chain")
	//env := config.G_CONFIG.Env
	list := controller.tokenService.FindList()
	utils.Success(ctx, list)
}

func (controller *tokenController) FindById(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *tokenController) Insert(ctx *gin.Context) {
	var dto dtos.InsertTokenDTO
	err := ctx.ShouldBind(&dto)
	if err != nil {
		utils.Fail(ctx, 500, err.Error())
	} else {
		token := models.Token{
			Name:      dto.Name,
			Symbol:    dto.Symbol,
			Logo:      dto.Logo,
			Decimals:  dto.Decimals,
			Address:   dto.Address,
			Balance:   "0",
			ChainType: dto.ChainType,
			Network:   dto.Network,
		}
		controller.tokenService.Insert(token)
		utils.Success(ctx)
	}
}

func (controller *tokenController) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *tokenController) Del(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
