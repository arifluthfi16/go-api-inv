package routers

import (
	"github.com/gin-gonic/gin"
	"invest/app/controllers"
)
type Crowdfund struct {}

func (crowdfund *Crowdfund) Route (route *gin.Engine){
	router := route.Group("/crowdfunding")
	Controller := controllers.CrowdfundController{}
	//Auth := controllers.AuthController{}
	//, Auth.VerifyAccess()
	//, Auth.VerifyAccess()
	//, Auth.VerifyAccess()
	//, Auth.VerifyAccess()
	//, Auth.VerifyAccess()

	router.POST("/", Controller.Create)
	router.GET("/:id", Controller.FindOneById)
	router.GET("/" ,Controller.FindAll)
	router.DELETE("/:id", Controller.DeleteById)
	router.PUT("/:id", Controller.Update)
}