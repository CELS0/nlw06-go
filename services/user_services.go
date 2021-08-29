package services

import (
	"github.com/CELS0/nlw06-go/models"
	repository "github.com/CELS0/nlw06-go/repositories"
)

func Create(p *models.User) {
	repository.NewCreate(p)
}

func Get(newid int) models.User {
	var p = repository.FindOne(newid)
	return p
}

func FindAll() []models.User {
	var p = repository.FindAll()
	return p
}
