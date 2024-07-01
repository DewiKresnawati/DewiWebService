package handlers

import (
	"github.com/DewiKresnawati/DewiWebService/database"
	"github.com/DewiKresnawati/DewiWebService/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateOrder handles creating a new order.
// @Summary Create a new order
// @Description Create a new order
// @Tags Orders
// @Accept json
// @Produce json
// @Param   order body models.OrderRequest true "Order data"
// @Success 201 {object} models.OrderResponse
// @Failure 400 {object} map[string]interface{}
// @Router /orders [post]
// @Security BearerAuth
func CreateOrder(c *fiber.Ctx) error {
	db := database.DB
	var req models.OrderRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	order := models.Order{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Total:     req.Total,
	}

	if err := db.Create(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	orderResponse := models.OrderResponse{
		ID:        order.ID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Total:     order.Total,
	}

	return c.Status(fiber.StatusCreated).JSON(orderResponse)
}

// GetAllOrders handles retrieving all orders.
// @Summary Get all orders
// @Description Retrieve all orders
// @Tags Orders
// @Accept json
// @Produce json
// @Success 200 {array} models.OrderResponse
// @Failure 500 {object} map[string]interface{}
// @Router /orders [get]
// @Security BearerAuth
func GetAllOrders(c *fiber.Ctx) error {
	db := database.DB
	var orders []models.Order
	if err := db.Find(&orders).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	orderResponses := make([]models.OrderResponse, 0, len(orders))
	for _, order := range orders {
		orderResponse := models.OrderResponse{
			ID:        order.ID,
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
			Total:     order.Total,
		}
		orderResponses = append(orderResponses, orderResponse)
	}

	return c.JSON(orderResponses)
}

// GetOrderByID handles retrieving an order by its ID.
// @Summary Get order by ID
// @Description Retrieve an order by its ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} models.OrderResponse
// @Failure 404 {object} map[string]interface{}
// @Router /orders/{id} [get]
// @Security BearerAuth
func GetOrderByID(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var order models.Order
	if err := db.First(&order, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	orderResponse := models.OrderResponse{
		ID:        order.ID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Total:     order.Total,
	}

	return c.JSON(orderResponse)
}

// UpdateOrder handles updating an existing order.
// @Summary Update order
// @Description Update an existing order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param   order body models.OrderRequest true "Order data"
// @Success 200 {object} models.OrderResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /orders/{id} [put]
// @Security BearerAuth
func UpdateOrder(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var req models.OrderRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var order models.Order
	if err := db.First(&order, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	order.ProductID = req.ProductID
	order.Quantity = req.Quantity
	order.Total = req.Total

	if err := db.Save(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	orderResponse := models.OrderResponse{
		ID:        order.ID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Total:     order.Total,
	}

	return c.JSON(orderResponse)
}

// DeleteOrder handles deleting an order.
// @Summary Delete order
// @Description Delete an order by its ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 204 {object} nil
// @Failure 500 {object} map[string]interface{}
// @Router /orders/{id} [delete]
// @Security BearerAuth
func DeleteOrder(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	if err := db.Delete(&models.Order{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
