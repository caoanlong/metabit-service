package main

import (
	"metabit-service/app/router"
	"metabit-service/core/config"
)

func main() {
	r := router.Routers()
	r.Run(config.G_CONFIG.Port)
}
