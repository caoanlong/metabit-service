package main

import (
	"metabit-service/admin/router"
	"metabit-service/core/config"
)

func main() {
	r := router.Routers()
	r.Run(config.G_CONFIG.Port)
}
