package controllers

import (
	"strconv"

	"github.com/CELS0/nlw06-go/models"
	"github.com/CELS0/nlw06-go/services"
	"github.com/gin-gonic/gin"
)

func TagCreateController(c *gin.Context) {
	var p *models.Tag
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	services.CreateTag(p)

	c.JSON(201, p)
}

func TagGetController(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var p = services.GetTag(newid)

	c.JSON(200, p)
}

func TagFindAllControllers(c *gin.Context) {
	var p = services.FindAllTag()
	c.JSON(200, p)
}
func TagDeleteController(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	services.DeleteTag(newid)

	c.Status(200)
}
func TagUpdateController(c *gin.Context) {
	var p *models.Tag

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	services.UpdateTag(p)

	c.JSON(200, p)
}
