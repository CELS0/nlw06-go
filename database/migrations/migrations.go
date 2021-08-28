package migrations

import (
	"github.com/CELS0/nlw06-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})

}
