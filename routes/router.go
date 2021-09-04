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
			user.POST("/", controllers.UserCreateController)
			user.GET("/:id", controllers.UserGetController)
			user.GET("/", controllers.UserFindAllControllers)
			user.DELETE("/:id", controllers.UserDeleteController)
			user.PUT("/", controllers.UserUpdateController)
		}
		tag := main.Group("tag")
		{
			tag.POST("/", controllers.TagCreateController)
			tag.GET("/:id", controllers.TagGetController)
			tag.GET("/", controllers.TagFindAllControllers)
			tag.DELETE("/:id", controllers.TagDeleteController)
			tag.PUT("/", controllers.TagUpdateController)
		}
	}
	return router
}
