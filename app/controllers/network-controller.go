package controllers

import (
	"github.com/gin-gonic/gin"
	"metabit-service/app/dtos"
	"metabit-service/core/models"
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

func (controller *networkController) FindAll(ctx *gin.Context) {
	err, list := controller.networkService.FindAll()
	if err != nil {
		utils.Fail(ctx, 500, err.Error())
		return
	}
	utils.Success(ctx, list)
}

func (controller *networkController) FindList(ctx *gin.Context) {
	//chain := ctx.Query("chain")
	//env := config.G_CONFIG.Env
	err, list := controller.networkService.FindList()
	if err != nil {
		utils.Fail(ctx, 500, err.Error())
		return
	}
	utils.Success(ctx, list)
}

func (controller *networkController) FindById(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *networkController) Insert(ctx *gin.Context) {
	var dto dtos.InsertNetworkDTO
	err := ctx.ShouldBind(&dto)
	if err != nil {
		utils.Fail(ctx, 400, err.Error())
		return
	}
	network := models.Network{
		Name:        dto.Name,
		ShortName:   dto.ShortName,
		NetworkType: dto.NetworkType,
		ChainType:   dto.ChainType,
		Api:         dto.Api,
		ScanUrl:     dto.ScanUrl,
	}
	err, _ = controller.networkService.Insert(&network)
	if err != nil {
		utils.Fail(ctx, 500, err.Error())
		return
	}
	utils.Success(ctx)
}

func (controller *networkController) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *networkController) Del(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
