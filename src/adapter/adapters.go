package adapter

import (
	orderFactory "order-api/src/app/order/factory"
	orderHandler "order-api/src/app/order/handler"
	"order-api/src/database"
)

type handlers struct {
	Order orderHandler.Handler
}

func Init() handlers {
	db := database.GetPosgreSQLConnection()
	return handlers{
		Order: orderFactory.Factory(db),
	}
}
