package routes

import (
	"github.com/CELS0/nlw06-go/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.UserController)
			user.GET("/:id", controllers.UserGetController)
			user.GET("/", controllers.UserFindAllControllers)
		}
	}
	return router
}
