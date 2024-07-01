package handlers

import (
	"github.com/DewiKresnawati/DewiWebService/database"
	"github.com/DewiKresnawati/DewiWebService/models"
	"github.com/gofiber/fiber/v2"
)

// CreateSupplier handles creating a new supplier.
// @Summary Create a new supplier
// @Description Create a new supplier
// @Tags Supplier
// @Accept json
// @Produce json
// @Param   supplier body models.SupplierRequest true "Supplier data"
// @Success 201 {object} models.SupplierResponse
// @Failure 400 {object} map[string]interface{}
// @Router /suppliers [post]
// @Security BearerAuth
func CreateSupplier(c *fiber.Ctx) error {
	db := database.DB
	var req models.SupplierRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	supplier := models.Supplier{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := db.Create(&supplier).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(supplier)
}

// GetAllSuppliers handles retrieving all suppliers.
// @Summary Get all suppliers
// @Description Retrieve all suppliers
// @Tags Supplier
// @Accept json
// @Produce json
// @Success 200 {array} models.SupplierResponse
// @Failure 500 {object} map[string]interface{}
// @Router /suppliers [get]
// @Security BearerAuth
func GetAllSuppliers(c *fiber.Ctx) error {
	db := database.DB
	var suppliers []models.Supplier
	if err := db.Find(&suppliers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(suppliers)
}

// GetSupplierByID handles retrieving a supplier by its ID.
// @Summary Get supplier by ID
// @Description Retrieve a supplier by its ID
// @Tags Supplier
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID"
// @Success 200 {object} models.SupplierResponse
// @Failure 404 {object} map[string]interface{}
// @Router /suppliers/{id} [get]
// @Security BearerAuth
func GetSupplierByID(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var supplier models.Supplier
	if err := db.First(&supplier, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Supplier not found",
		})
	}

	return c.JSON(supplier)
}

// UpdateSupplier handles updating an existing supplier.
// @Summary Update supplier
// @Description Update an existing supplier
// @Tags Supplier
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID"
// @Param   supplier body models.SupplierRequest true "Supplier data"
// @Success 200 {object} models.SupplierResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /suppliers/{id} [put]
// @Security BearerAuth
func UpdateSupplier(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var req models.SupplierRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var supplier models.Supplier
	if err := db.First(&supplier, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Supplier not found",
		})
	}

	supplier.Name = req.Name
	supplier.Email = req.Email

	if err := db.Save(&supplier).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(supplier)
}

// DeleteSupplier handles deleting a supplier.
// @Summary Delete supplier
// @Description Delete a supplier by its ID
// @Tags Supplier
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID"
// @Success 204 {object} nil
// @Failure 500 {object} map[string]interface{}
// @Router /suppliers/{id} [delete]
// @Security BearerAuth
func DeleteSupplier(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	if err := db.Delete(&models.Supplier{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
