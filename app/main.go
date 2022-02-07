package main

import (
	"github.com/gin-gonic/gin"
	"metabit-service/api/config"
	_ "metabit-service/api/config"
	"metabit-service/app/controllers"
)

var (
	tokenController = controllers.NewTokenController()
)

func main() {
	router := gin.Default()

	tokens := router.Group("/tokens")
	{
		tokens.GET("", tokenController.FindList)
		tokens.GET("/:id", tokenController.FindById)
		tokens.POST("", tokenController.Insert)
		tokens.PUT("/:id", tokenController.Update)
		tokens.DELETE("/:id", tokenController.Del)
	}

	router.Run(config.G_CONFIG.Port)
}
