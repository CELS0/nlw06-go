package repository

import (
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/CELS0/nlw06-go/database"
	"github.com/CELS0/nlw06-go/models"
)

func sHA256Encoder(s string) string {
	str := sha256.Sum256([]byte(s))

	return fmt.Sprintf("%x", str)
}

func NewCreate(p *models.User) {
	db := database.GetDatabase()

	p.Password = sHA256Encoder(p.Password)

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
