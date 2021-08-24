package routers

import (
	"github.com/gin-gonic/gin"
	"invest/app/controllers"
)

type RequestKYC struct {}

func (req *RequestKYC) Route(route *gin.Engine) {
	router := route.Group("/kyc")
	Controller := controllers.RequestKYCController{}
	router.POST("/", Controller.Create)
	router.GET("/:id", Controller.FindOneById)
	router.DELETE("/:id", Controller.DeleteById)
	router.GET("/", Controller.FindAllKYCRequest)
	router.POST("/approve", Controller.ApproveValidation)
	router.PUT("/:id", Controller.Update)
}