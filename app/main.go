package main

import (
	"metabit-service/app/router"
	"metabit-service/core/config"
)

func main() {
	config.Init()
	r := router.Routers()
	r.Run(config.GConfig.Port)
}
