package main

import (
	"order-api/src/config/env"
	"order-api/src/router"
)

func main() {
	router.StartServer().Run(":" + env.GetServerEnv())
}
