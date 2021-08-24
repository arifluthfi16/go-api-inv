package routers

import (
	"github.com/gin-gonic/gin"
	"invest/app/controllers"
)

type User struct {
}

func (user *User) Route (route *gin.Engine){
	router := route.Group("/user")
	Auth := controllers.AuthController{}

	router.POST("/", controllers.CreateUser)
	router.POST("/login", Auth.Login)
	router.GET("/:userid", controllers.FindUserById)
	router.GET("/", controllers.FindAllUser)
	router.DELETE("/:userid", controllers.DeleteUserById)
	router.PUT("/:userid", controllers.UpdateUser)
}
