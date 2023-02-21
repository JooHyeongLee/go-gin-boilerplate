package server

import (
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	//router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("member")
		{
			user := new(controllers.MemberController)
			userGroup.GET("/:id", user.Retrieve)
		}
	}
	return router

}
