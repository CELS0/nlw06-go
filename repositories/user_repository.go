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
		log.Fatal("cannot create user: " + err.Error())
	}
}

func FindOne(newid int) models.User {
	db := database.GetDatabase()
	var p models.User

	err := db.First(&p, newid).Error
	if err != nil {
		log.Fatal("cannot find user: " + err.Error())
	}

	return p
}

func FindAll() []models.User {
	db := database.GetDatabase()

	var p []models.User
	err := db.Find(&p).Error
	if err != nil {
		log.Fatal("cannot list users: " + err.Error())
	}
	return p
}

func Delete(newid int) {
	db := database.GetDatabase()
	err := db.Delete(&models.User{}, newid).Error
	if err != nil {
		log.Fatal("cannot delete user: " + err.Error())
	}
}

func Update(p *models.User) {
	db := database.GetDatabase()

	err := db.Save(&p).Error
	if err != nil {
		log.Fatal("cannot update user: " + err.Error())
	}
}
