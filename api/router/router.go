package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iquzart/go-rest-api/api/controller"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.Default()

	// Route v1
	v1 := r.Group("/v1")
	{
		v1.GET("/", controller.Home)
		v1.POST("/postjson", controller.PostHome)
		v1.GET("/querystring", controller.QueryString)
		v1.GET("/path/:name/:age", controller.PathParameters)

	}

	r.NoRoute(controller.NoFound)

	return r
}
