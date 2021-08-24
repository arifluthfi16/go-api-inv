package routers

import (
	"github.com/gin-gonic/gin"
	"invest/app/controllers"
)

type Admin struct {}

func (admin *Admin) Route(route *gin.Engine) {
	Auth := controllers.AuthController{}

	router := route.Group("/admin")
	Controller := controllers.AdminController{}
	router.POST("/", Controller.CreateAdmin)
	router.POST("/login", Auth.AdminLogin)
}