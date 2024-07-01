package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
}

type RegisterRequest struct {
	Username string
	Password string
}

type RegisterResponse struct {
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Token string
}
