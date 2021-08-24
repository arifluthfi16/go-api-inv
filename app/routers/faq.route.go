package routers

import (
	"github.com/gin-gonic/gin"
	"invest/app/controllers"
)
type FAQ struct {}

func (faq *FAQ) Route (route *gin.Engine){
	router := route.Group("/faq")
	Controller := controllers.FAQController{}

	router.POST("/", Controller.Create)
	router.GET("/:faq_id", Controller.FindOneById)
	router.GET("/", Controller.FindAll)
	router.DELETE("/:faq_id", Controller.DeleteById)
	router.PUT("/:faq_id", Controller.Update)
}

