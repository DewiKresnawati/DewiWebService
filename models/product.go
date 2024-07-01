package models

import "gorm.io/gorm"

// Product represents a product entity.
type Product struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Price       float64  `gorm:"not null"`
	CategoryID  uint     `gorm:"not null"`
	Category    Category // Relasi belongs to
	SupplierID  uint     `gorm:"not null"`
	Supplier    Supplier // Relasi belongs to
}

type ProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  uint    `json:"category_id"`
	SupplierID  uint    `json:"supplier_id"`
}

type ProductResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  uint    `json:"category_id"`
	SupplierID  uint    `json:"supplier_id"`
}
