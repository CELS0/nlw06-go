package controllers

import (
	"github.com/CELS0/nlw06-go/models"
	"github.com/CELS0/nlw06-go/services"
	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {
	var p models.User
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
