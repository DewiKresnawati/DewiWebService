package migration

import (
	"fmt"
	"log"

	"github.com/DewiKresnawati/DewiWebService/database"
	"github.com/DewiKresnawati/DewiWebService/models"
)

// RunMigration migrates the database schema.
func RunMigration() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Auto-migrate models
	err = db.AutoMigrate(&models.Product{}, &models.Category{}, &models.Order{}, &models.Supplier{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migrasi berhasil dijalankan")
}
