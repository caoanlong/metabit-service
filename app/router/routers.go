package router

import (
	"github.com/gin-gonic/gin"
	"metabit-service/app/controllers"
)

var (
	networkController = controllers.NewNetworkController()
	tokenController   = controllers.NewTokenController()
)

func Routers() *gin.Engine {
	router := gin.Default()

	networks := router.Group("/networks")
	{
		networks.GET("", networkController.FindAll)
		networks.GET("/:id", networkController.FindById)
		networks.POST("", networkController.Insert)
		networks.PUT("/:id", networkController.Update)
		networks.DELETE("/:id", networkController.Del)
	}

	tokens := router.Group("/tokens")
	{
		tokens.GET("", tokenController.FindAll)
		tokens.GET("/:id", tokenController.FindById)
		tokens.POST("", tokenController.Insert)
		tokens.PUT("/:id", tokenController.Update)
		tokens.DELETE("/:id", tokenController.Del)
	}
	return router
}
