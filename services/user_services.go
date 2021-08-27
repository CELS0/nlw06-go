package services

import (
	"github.com/CELS0/nlw06-go/models"
	repository "github.com/CELS0/nlw06-go/repositories"
)

func Create(p models.User) {
	repository.NewCreate(p)
}
