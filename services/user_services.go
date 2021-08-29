package services

import (
	"github.com/CELS0/nlw06-go/models"
	repository "github.com/CELS0/nlw06-go/repositories"
)

func CreateUser(p *models.User) {
	repository.NewCreate(p)
}

func GetUser(newid int) models.User {
	var p = repository.FindOne(newid)
	return p
}

func FindAllUser() []models.User {
	var p = repository.FindAll()
	return p
}

func DeleteUser(newid int) {
	repository.Delete(newid)
}

func UpdateUser(p *models.User) {
	repository.Update(p)
}
