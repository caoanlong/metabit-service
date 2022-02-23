package router

import (
	"github.com/gin-gonic/gin"
	"metabit-service/admin/controllers"
)

var (
	tokenController   = controllers.NewTokenController()
	networkController = controllers.NewNetworkController()
	sysUserController = controllers.NewSysUserController()
)

func Routers() *gin.Engine {
	router := gin.Default()
	admin := router.Group("/admin")

	tokens := admin.Group("/tokens")
	{
		tokens.GET("", tokenController.FindAll)
		tokens.GET("/:id", tokenController.FindById)
		tokens.POST("", tokenController.Insert)
		tokens.PUT("/:id", tokenController.Update)
		tokens.DELETE("/:id", tokenController.Del)
	}
	networks := admin.Group("/networks")
	{
		networks.GET("", networkController.FindAll)
		networks.GET("/:id", networkController.FindById)
		networks.POST("", networkController.Insert)
		networks.PUT("/:id", networkController.Update)
		networks.DELETE("/:id", networkController.Del)
	}
	sysusers := admin.Group("/sysusers")
	{
		sysusers.POST("/login", sysUserController.Login)
		sysusers.GET("", sysUserController.FindAll)
		sysusers.GET("/:id", sysUserController.FindById)
		sysusers.POST("", sysUserController.Insert)
		sysusers.PUT("/:id", sysUserController.Update)
		sysusers.DELETE("/:id", sysUserController.Del)
	}

	return router
}
