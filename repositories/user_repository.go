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
