package database

import (
	"fmt"
	"log"
	order "order-api/src/app/order/repository/record"
	"order-api/src/config/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresConn *gorm.DB

func InitPosgreSQL() {
	var err error

	config := env.GetSQLEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DBName, config.Port)
	postgresConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	postgresConn.AutoMigrate(&order.Order{}, &order.Item{})
}

func GetPosgreSQLConnection() *gorm.DB {
	return postgresConn
}
