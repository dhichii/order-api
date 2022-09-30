package factory

import (
	"order-api/src/app/order/handler"
	"order-api/src/app/order/repository"
	"order-api/src/app/order/service"

	"gorm.io/gorm"
)

func Factory(db *gorm.DB) handler.Handler {
	repo := repository.NewGORMRepository(db)
	serv := service.NewService(repo)
	return *handler.NewHandler(serv)
}
