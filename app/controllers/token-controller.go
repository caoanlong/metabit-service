package controllers

import (
	"github.com/gin-gonic/gin"
	"metabit-service/api/config"
	"metabit-service/api/models"
	"metabit-service/api/services"
	"metabit-service/api/utils"
	"metabit-service/app/dtos"
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
	utils.Success(ctx)
}

func (controller *tokenController) FindList(ctx *gin.Context) {
	//chain := ctx.Query("chain")
	env := config.G_CONFIG.Env
	utils.Success(ctx, env)
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
			Name:     dto.Name,
			Symbol:   dto.Symbol,
			Logo:     dto.Logo,
			Decimals: dto.Decimals,
			Address:  dto.Address,
			Balance:  "0",
			Type:     dto.Type,
			Category: dto.Category,
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