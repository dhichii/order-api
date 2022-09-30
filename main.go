package main

import (
	"order-api/src/config/env"
	"order-api/src/database"
	"order-api/src/router"
)

func main() {
	database.InitPosgreSQL()
	router.StartServer().Run(":" + env.GetServerEnv())
}
