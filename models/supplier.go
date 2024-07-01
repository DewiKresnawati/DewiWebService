package models

import "gorm.io/gorm"

// Supplier represents a supplier entity.
type Supplier struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Email string `gorm:"unique;not null"`
}

type SupplierRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SupplierResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
