package main

import (
	"metabit-service/api/config"
	"metabit-service/app/router"
)

func main() {
	r := router.Routers()
	r.Run(config.G_CONFIG.Port)
}
