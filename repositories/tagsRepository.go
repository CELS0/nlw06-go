package repository

import (
	"log"

	"github.com/CELS0/nlw06-go/database"
	"github.com/CELS0/nlw06-go/models"
)

func CreateTag(p *models.Tag) {
	db := database.GetDatabase()

	err := db.Create(&p).Error
	if err != nil {
		log.Fatal("cannot create Tag: " + err.Error())
	}
}

func FindOneTag(newid int) models.Tag {
	db := database.GetDatabase()
	var p models.Tag

	err := db.First(&p, newid).Error
	if err != nil {
		log.Fatal("cannot find Tag: " + err.Error())
	}

	return p
}

func FindAllTag() []models.Tag {
	db := database.GetDatabase()

	var p []models.Tag
	err := db.Find(&p).Error
	if err != nil {
		log.Fatal("cannot list Tags: " + err.Error())
	}
	return p
}

func DeleteTag(newid int) {
	db := database.GetDatabase()
	err := db.Delete(&models.Tag{}, newid).Error
	if err != nil {
		log.Fatal("cannot delete Tag: " + err.Error())
	}
}

func UpdateTag(p *models.Tag) {
	db := database.GetDatabase()

	err := db.Save(&p).Error
	if err != nil {
		log.Fatal("cannot update Tag: " + err.Error())
	}
}
