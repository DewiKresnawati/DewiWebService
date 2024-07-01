package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID uint    `gorm:"not null"`
	Product   Product // Contoh: Relasi has many, sesuaikan dengan struktur Product Anda
	Quantity  uint    `gorm:"not null"`
	Total     float64 `gorm:"not null"`
}

type OrderRequest struct {
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Total     float64 `json:"total"`
}

type OrderResponse struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Total     float64 `json:"total"`
}
