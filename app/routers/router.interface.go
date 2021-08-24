package routers

import "github.com/gin-gonic/gin"

type RouteGroup interface {
	Route(*gin.Engine)
}

