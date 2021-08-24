package routers

import (
	"github.com/gin-gonic/gin"
	"invest/app/controllers"
)

type Investor struct {}

func (investor *Investor) Route (route *gin.Engine){
	router := route.Group("/investor")
	Controller := controllers.InvestorController{}

	router.POST("/", Controller.Create)
	router.GET("/:id", Controller.FindOneById)
	router.GET("/", Controller.FindAll)
	router.GET("/crowdfund/:id", Controller.FindAll)
	router.DELETE("/:id", Controller.DeleteById)
	router.PUT("/:id", Controller.Update)
}

