package model

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
