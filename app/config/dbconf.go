package config

type DBConf struct{
	DBHost string
	DBUser string
	DBName string
	DBPort string
	DbPass string
}

var DBConfig DBConf