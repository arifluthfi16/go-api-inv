package routers

import (
	"github.com/gin-gonic/gin"
	"invest/app/controllers"
)

type Transparansi struct {}

func (transparansi *Transparansi) Route(route *gin.Engine) {
	router := route.Group("/transparansi")
	Controller := controllers.TransparansiController{}

	router.POST("/", Controller.Create)
	router.GET("/:id", Controller.FindOneById)
	router.GET("/", Controller.FindAll)
	router.DELETE("/:id", Controller.DeleteById)
	router.PUT("/:id", Controller.Update)
}