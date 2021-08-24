package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Loaders struct{}

func (l Loaders) getEnv(key, fallback string) string{
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (l Loaders) LoadEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	AppConfig.AppName 	= l.getEnv("APP_NAME", "VanameID")
	AppConfig.AppEnv 	= l.getEnv("APP_ENV", "development")
	AppConfig.AppPort 	= l.getEnv("APP_PORT", "4040")
	AppConfig.JwtKey 	= l.getEnv("JWT_KEY", "12345678")

	DBConfig.DBHost 	= l.getEnv("DB_HOST", "localhost")
	DBConfig.DBName 	= l.getEnv("DB_NAME", "gotoko")
	DBConfig.DBUser 	= l.getEnv("DB_USER", "postgres")
	DBConfig.DbPass 	= l.getEnv("DB_PASSWORD", "root")
	DBConfig.DBPort 	= l.getEnv("DB_PORT", "5432")
}
