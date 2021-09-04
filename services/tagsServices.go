package services

import (
	"github.com/CELS0/nlw06-go/models"
	repository "github.com/CELS0/nlw06-go/repositories"
)

func CreateTag(p *models.Tag) {
	repository.CreateTag(p)
}

func GetTag(newid int) models.Tag {
	var p = repository.FindOneTag(newid)
	return p
}

func FindAllTag() []models.Tag {
	var p = repository.FindAllTag()
	return p
}

func DeleteTag(newid int) {
	repository.DeleteTag(newid)
}

func UpdateTag(p *models.Tag) {
	repository.UpdateTag(p)
}
