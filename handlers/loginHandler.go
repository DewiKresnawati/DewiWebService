package handlers

import (
	"github.com/DewiKresnawati/DewiWebService/database"
	"github.com/DewiKresnawati/DewiWebService/models"
	"github.com/DewiKresnawati/DewiWebService/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Login
// @Description Login with username and password
// @ID login
// @Accept  json
// @Produce  json
// @Param   login  body     models.LoginRequest  true  "Login Request"
// @Success 200    {object} map[string]interface{}
// @Failure 400    {object} map[string]interface{}
// @Failure 401    {object} map[string]interface{}
// @Failure 500    {object} map[string]interface{}
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	db := database.DB
	loginRequest := new(models.LoginRequest)
	err := c.BodyParser(loginRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad request",
			"error":   err.Error(),
		})
	}

	user := new(models.User)
	err = db.Where("username = ?", loginRequest.Username).First(user).Error
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "username belum terdaftar",
			"error":   err.Error(),
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "password tidak sesuai",
			"error":   err.Error(),
		})
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate token",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login successful",
		"token":   token,
	})
}