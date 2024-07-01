package handlers

import (
	"github.com/DewiKresnawati/DewiWebService/database"
	"github.com/DewiKresnawati/DewiWebService/models"
	"github.com/DewiKresnawati/DewiWebService/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Register
// @Description Register a new user
// @ID register
// @Accept  json
// @Produce  json
// @Param   register  body     models.RegisterRequest  true  "Register Request"
// @Success 201    {object} map[string]interface{}
// @Failure 400    {object} map[string]interface{}
// @Failure 500    {object} map[string]interface{}
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to hash password",
			"error":   err.Error(),
		})
	}

	user.Password = string(hashedPassword)

	result := db.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to register user",
			"error":   result.Error.Error(),
		})
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate token",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user registered successfully",
		"token":   token,
	})
}
