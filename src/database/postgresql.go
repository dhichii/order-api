package database

import (
	"fmt"
	"log"
	"order-api/src/config/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPosgreSQL() *gorm.DB {
	config := env.GetSQLEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DBName, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	return db
}
