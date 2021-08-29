package controllers

import (
	"strconv"

	"github.com/CELS0/nlw06-go/models"
	"github.com/CELS0/nlw06-go/services"
	"github.com/gin-gonic/gin"
)

func UserCreateController(c *gin.Context) {
	var p *models.User
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	services.CreateUser(p)

	c.JSON(201, p)
}

func UserGetController(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var p = services.GetUser(newid)

	c.JSON(200, p)
}

func UserFindAllControllers(c *gin.Context) {
	var p = services.FindAllUser()
	c.JSON(200, p)
}
func UserDeleteController(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	services.DeleteUser(newid)

	c.Status(200)
}
func UserUpdateController(c *gin.Context) {
	var p *models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	services.UpdateUser(p)

	c.JSON(200, p)
}
