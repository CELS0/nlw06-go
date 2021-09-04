package repository

import (
	"log"

	"github.com/CELS0/nlw06-go/database"
	"github.com/CELS0/nlw06-go/models"
)

func CreateUser(p *models.User) {
	db := database.GetDatabase()

	err := db.Create(&p).Error
	if err != nil {
		log.Fatal("cannot create user: " + err.Error())
	}
}

func FindOneUser(newid int) models.User {
	db := database.GetDatabase()
	var p models.User

	err := db.First(&p, newid).Error
	if err != nil {
		log.Fatal("cannot find user: " + err.Error())
	}

	return p
}

func FindAllUser() []models.User {
	db := database.GetDatabase()

	var p []models.User
	err := db.Find(&p).Error
	if err != nil {
		log.Fatal("cannot list users: " + err.Error())
	}
	return p
}

func DeleteUser(newid int) {
	db := database.GetDatabase()
	err := db.Delete(&models.User{}, newid).Error
	if err != nil {
		log.Fatal("cannot delete user: " + err.Error())
	}
}

func UpdateUser(p *models.User) {
	db := database.GetDatabase()

	err := db.Save(&p).Error
	if err != nil {
		log.Fatal("cannot update user: " + err.Error())
	}
}
