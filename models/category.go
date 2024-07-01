package models

import "gorm.io/gorm"

// Category represents a category entity.
type Category struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

type CategoryRequest struct {
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
