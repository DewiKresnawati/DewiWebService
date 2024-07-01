package handlers

import (
	"errors"

	"github.com/DewiKresnawati/DewiWebService/database"
	"github.com/DewiKresnawati/DewiWebService/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary Create a new product
// @Description Create a new product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.ProductRequest true "Product data"
// @Success 201 {object} models.ProductResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products [post]
// @Security BearerAuth
// @TokenUrl http://localhost:4111/api/v1/login
func CreateProduct(c *fiber.Ctx) error {
	// Parse request body into ProductRequest struct
	var req models.ProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	// Create a new Product instance
	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
		SupplierID:  req.SupplierID,
	}

	// Get the database connection
	db := database.DB

	// Create the product in the database
	if err := db.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create product",
			"error":   err.Error(),
		})
	}

	// Return the created product as response
	return c.Status(fiber.StatusCreated).JSON(product)
}

// @Summary Get all products
// @Description Get all products
// @Tags Products
// @Produce json
// @Success 200 {array} models.ProductResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products [get]
// @Security BearerAuth
func GetAllProducts(c *fiber.Ctx) error {
	// Get the database connection
	db := database.DB

	// Query all products from the database
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
		})
	}

	// Return the products as response
	return c.JSON(products)
}

// @Summary Get product by ID
// @Description Get product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.ProductResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{id} [get]
// @Security BearerAuth
func GetProductByID(c *fiber.Ctx) error {
	// Get the product ID from the URL parameters
	id := c.Params("id")

	// Get the database connection
	db := database.DB

	// Query the product from the database by ID
	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Product not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch product",
			"error":   err.Error(),
		})
	}

	// Return the product as response
	return c.JSON(product)
}

// @Summary Update product by ID
// @Description Update product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.ProductRequest true "Product data"
// @Success 200 {object} models.ProductResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{id} [put]
// @Security BearerAuth
func UpdateProduct(c *fiber.Ctx) error {
	// Get the product ID from the URL parameters
	id := c.Params("id")

	// Parse request body into ProductRequest struct
	var req models.ProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	// Get the database connection
	db := database.DB

	// Query the product from the database by ID
	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Product not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch product",
			"error":   err.Error(),
		})
	}

	// Update the product fields
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.CategoryID = req.CategoryID
	product.SupplierID = req.SupplierID

	// Save the updated product to the database
	if err := db.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update product",
			"error":   err.Error(),
		})
	}

	// Return the updated product as response
	return c.JSON(product)
}

// @Summary Delete product by ID
// @Description Delete product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 204 {object} nil
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{id} [delete]
// @Security BearerAuth
func DeleteProduct(c *fiber.Ctx) error {
	// Get the product ID from the URL parameters
	id := c.Params("id")

	// Get the database connection
	db := database.DB

	// Delete the product from the database by ID
	if err := db.Delete(&models.Product{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete product",
			"error":   err.Error(),
		})
	}

	// Return success message
	return c.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
