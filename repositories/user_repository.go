package repository

import (
	"log"

	"github.com/CELS0/nlw06-go/database"
	"github.com/CELS0/nlw06-go/models"
)

func NewCreate(p *models.User) {
	db := database.GetDatabase()
	log.Print(p.ID)
	err := db.Create(&p).Error
	if err != nil {
		log.Fatal("error ")
	}
}

func FindOne(newid int) models.User {
	db := database.GetDatabase()
	var p models.User

	err := db.First(&p, newid).Error
	if err != nil {
		log.Fatal("error ")
	}

	return p
}

func FindAll() []models.User {
	db := database.GetDatabase()

	var p []models.User
	err := db.Find(&p).Error
	if err != nil {
		log.Fatal("error")
	}
	return p
}
