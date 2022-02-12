package router

import (
	"github.com/gin-gonic/gin"
	"metabit-service/app/controllers"
)

var (
	tokenController = controllers.NewTokenController()
)

func Routers() *gin.Engine {
	router := gin.Default()

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
