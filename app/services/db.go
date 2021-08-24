package services

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	c "invest/app/config"
)

type DBManager struct {
	DB *gorm.DB
}

var Manager DBManager

func LoadDB(){
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		c.DBConfig.DBHost, c.DBConfig.DBUser, c.DBConfig.DbPass, c.DBConfig.DBName, c.DBConfig.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	Manager.DB = db
}