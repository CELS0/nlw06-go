package repository

import (
	"log"

	"github.com/CELS0/nlw06-go/database"
	"github.com/CELS0/nlw06-go/models"
)

func NewCreate(p models.User) {
	log.Println("REPOSITORY", p)
	db := database.GetDatabase()

	err := db.Create(&p).Error
	if err != nil {
		log.Fatal("error ")
	}
}

// func FindOne(c *gin.Context) {
// 	id := c.Param("id")
// 	db := database.GetDatabase()
// 	var p models.User

// 	newid, err := strconv.Atoi(id)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "ID has to be integer",
// 		})
// 		return
// 	}
// 	err = db.First(&p, newid).Error
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "cannot find user:" + err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(200, p)
// }

// func FindAll(c *gin.Context) {
// 	db := database.GetDatabase()

// 	var ps []models.User
// 	err := db.Find(&ps).Error
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "cannot list users: " + err.Error(),
// 		})
// 	}
// 	c.JSON(200, ps)
// }
