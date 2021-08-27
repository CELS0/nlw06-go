package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Admin     bool      `json:"admin"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}
