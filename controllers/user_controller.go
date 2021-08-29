package controllers

import (
	"fmt"
	"strconv"

	"github.com/CELS0/nlw06-go/models"
	"github.com/CELS0/nlw06-go/services"
	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {
	var p *models.User
	fmt.Print(p)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	services.Create(p)

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

	var p = services.Get(newid)

	c.JSON(200, p)
}

func UserFindAllControllers(c *gin.Context) {
	var p = services.FindAll()
	c.JSON(200, p)
}
