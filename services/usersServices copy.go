package services

import (
	"crypto/sha256"
	"fmt"

	"github.com/CELS0/nlw06-go/models"
	repository "github.com/CELS0/nlw06-go/repositories"
)

func CreateUser(p *models.User) {
	p.Password = sHA256Encoder(p.Password)
	repository.CreateUser(p)
}

func GetUser(newid int) models.User {
	var p = repository.FindOneUser(newid)
	return p
}

func FindAllUser() []models.User {
	var p = repository.FindAllUser()
	return p
}

func DeleteUser(newid int) {
	repository.DeleteUser(newid)
}

func UpdateUser(p *models.User) {
	repository.UpdateUser(p)
}

func sHA256Encoder(s string) string {
	str := sha256.Sum256([]byte(s))

	return fmt.Sprintf("%x", str)
}
