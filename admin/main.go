package main

import (
	"github.com/gin-gonic/gin"
	"metabit-service/admin/controllers"
	"metabit-service/api/config"
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
	}

	router.Run(config.G_CONFIG.Port)
}
