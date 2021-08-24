package config

type AppConf struct {
	AppPort string
	AppName string
	AppEnv string
	JwtKey string
}

var AppConfig AppConf