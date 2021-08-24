package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"invest/app/config"
	"invest/app/models"
	"invest/app/routers"
	"invest/app/services"
	"invest/app/utils"
	"log"
)

type Server struct {
	Router *gin.Engine
}

func (server *Server) Initialize(){
	server.InitializeDefaultRoutes()
}

func (server *Server) InitializeDefaultRoutes(){
	var errorHandler = utils.ErrorHandler{}
	server.Router = gin.Default()
	server.Router.MaxMultipartMemory = 8 << 20
	server.Router.Use(errorHandler.JSONAppErrorReporter())
}

func (server *Server) dbMigrate(){
	var err error
	for _,model := range models.RegisterModels(){
		err = services.Manager.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database Migration succeed")
}

func (server *Server) loadRoute(){
	var Router = routers.IndexRouter{}
	for _, routes := range Router.LoadRoutes(){
		routes.Route(server.Router)
	}
}

func (server *Server) Run(appConfig config.AppConf){
	fmt.Println("Rise and shine! 🌞🌞🌞")
	fmt.Println("Listening on port : "+appConfig.AppPort)
	server.Router.Run("127.0.0.1:"+appConfig.AppPort)
}

func Run(){
	var server = Server{}
	server.Initialize()
	server.loadRoute()
	server.Run(config.AppConfig)
}

