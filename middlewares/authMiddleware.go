package middlewares

import (
    "strings"

    "github.com/DewiKresnawati/DewiWebService/utils"
    "github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "Unauthorized: No token provided",
            })
        }

        tokenParts := strings.Split(authHeader, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "Unauthorized: Invalid authorization header format",
            })
        }
        token := tokenParts[1]

        claims, err := utils.VerifyToken(token)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "Unauthorized: " + err.Error(),
            })
        }

        c.Locals("user", claims)

        return c.Next()
    }
}
