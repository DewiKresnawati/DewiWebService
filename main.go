package main

import (
	"log"

	"github.com/DewiKresnawati/DewiWebService/database/migration"
	_ "github.com/DewiKresnawati/DewiWebService/docs"
	"github.com/DewiKresnawati/DewiWebService/middlewares"
	"github.com/DewiKresnawati/DewiWebService/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger" // swagger handler
)

// @title Golang JWT Auth API
// @version 1.0
// @description This is a sample JWT auth server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4113
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	migration.RunMigration()

	app := fiber.New()
	app.Use(cors.New())

	// Route to Swagger docs
	app.Get("/swagger/*", swagger.HandlerDefault) // use more specific route for Swagger

	// Initialize other routes
	routes.RouteInit(app)

	// Example secure endpoint with JWT authentication
	app.Get("/api/v1/", middlewares.AuthMiddleware(), func(c *fiber.Ctx) error {
		// Token is valid, continue processing
		user := c.Locals("user")
		return c.JSON(fiber.Map{
			"message": "You are authorized!",
			"user":    user,
		})
	})

	err := app.Listen(":4113")
	if err != nil {
		log.Fatal(err)
	}
}
