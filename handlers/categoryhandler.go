package handlers

import (
    "github.com/DewiKresnawati/DewiWebService/database"
    "github.com/DewiKresnawati/DewiWebService/models"
    "github.com/gofiber/fiber/v2"
)

// CreateCategory handles creating a new category.
// @Summary Create a new category
// @Description Create a new category
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body models.CategoryRequest true "Category data"
// @Success 201 {object} models.CategoryResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /categories [post]
// @Security BearerAuth
func CreateCategory(c *fiber.Ctx) error {
    db := database.DB
    var req models.CategoryRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    category := models.Category{
        Name: req.Name,
    }

    if err := db.Create(&category).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    response := models.CategoryResponse{
        ID:   category.ID,
        Name: category.Name,
    }

    return c.Status(fiber.StatusCreated).JSON(response)
}

// GetAllCategories handles retrieving all categories.
// @Summary Get all categories
// @Description Retrieve all categories
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {array} models.CategoryResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /categories [get]
// @Security BearerAuth
func GetAllCategories(c *fiber.Ctx) error {
    db := database.DB
    var categories []models.Category
    if err := db.Find(&categories).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    var response []models.CategoryResponse
    for _, category := range categories {
        response = append(response, models.CategoryResponse{
            ID:   category.ID,
            Name: category.Name,
        })
    }

    return c.JSON(response)
}

// GetCategoryByID handles retrieving a category by its ID.
// @Summary Get category by ID
// @Description Retrieve a category by its ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.CategoryResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /categories/{id} [get]
// @Security BearerAuth
func GetCategoryByID(c *fiber.Ctx) error {
    db := database.DB
    id := c.Params("id")
    var category models.Category
    if err := db.First(&category, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Category not found",
        })
    }

    response := models.CategoryResponse{
        ID:   category.ID,
        Name: category.Name,
    }

    return c.JSON(response)
}

// UpdateCategory handles updating an existing category.
// @Summary Update category
// @Description Update an existing category
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body models.CategoryRequest true "Category data"
// @Success 200 {object} models.CategoryResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /categories/{id} [put]
// @Security BearerAuth
func UpdateCategory(c *fiber.Ctx) error {
    db := database.DB
    id := c.Params("id")
    var req models.CategoryRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    var category models.Category
    if err := db.First(&category, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Category not found",
        })
    }

    category.Name = req.Name

    if err := db.Save(&category).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    response := models.CategoryResponse{
        ID:   category.ID,
        Name: category.Name,
    }

    return c.JSON(response)
}

// DeleteCategory handles deleting a category.
// @Summary Delete category
// @Description Delete a category by its ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 204 {object} nil
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /categories/{id} [delete]
// @Security BearerAuth
func DeleteCategory(c *fiber.Ctx) error {
    db := database.DB
    id := c.Params("id")
    if err := db.Delete(&models.Category{}, id).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.SendStatus(fiber.StatusNoContent)
}
