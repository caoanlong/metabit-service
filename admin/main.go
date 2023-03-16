package main

import (
	"metabit-service/admin/router"
	"metabit-service/core/config"
)

func main() {
	config.Init()
	r := router.Routers()
	r.Run(config.GConfig.Port)
}
