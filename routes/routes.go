package routes

import (
	"github.com/DewiKresnawati/DewiWebService/handlers"
	"github.com/DewiKresnawati/DewiWebService/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	r := app.Group("/api/v1")
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome in webservice Dewi!")
	})

	// auth route
	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)
	r.Get("/protected", middlewares.AuthMiddleware(), handlers.ProtectedRoute)
	r.Post("/logout", middlewares.AuthMiddleware(), handlers.Logout)

	// Product routes
	r.Post("/products", middlewares.AuthMiddleware(), handlers.CreateProduct)
	r.Get("/products", middlewares.AuthMiddleware(), handlers.GetAllProducts)
	r.Get("/products/:id", middlewares.AuthMiddleware(), handlers.GetProductByID)
	r.Put("/products/:id", middlewares.AuthMiddleware(), handlers.UpdateProduct)
	r.Delete("/products/:id", handlers.DeleteProduct)

	// Category routes
	r.Post("/categories", middlewares.AuthMiddleware(), handlers.CreateCategory)
	r.Get("/categories", middlewares.AuthMiddleware(), handlers.GetAllCategories)
	r.Get("/categories/:id", middlewares.AuthMiddleware(), handlers.GetCategoryByID)
	r.Put("/categories/:id", middlewares.AuthMiddleware(), handlers.UpdateCategory)
	r.Delete("/categories/:id", handlers.DeleteCategory)

	// Order routes
	r.Post("/orders", middlewares.AuthMiddleware(), handlers.CreateOrder)
	r.Get("/orders", middlewares.AuthMiddleware(), handlers.GetAllOrders)
	r.Get("/orders/:id", middlewares.AuthMiddleware(), handlers.GetOrderByID)
	r.Put("/orders/:id", middlewares.AuthMiddleware(), handlers.UpdateOrder)
	r.Delete("/orders/:id", handlers.DeleteOrder)

	// Supplier routes
	r.Post("/suppliers", middlewares.AuthMiddleware(), handlers.CreateSupplier)
	r.Get("/suppliers", middlewares.AuthMiddleware(), handlers.GetAllSuppliers)
	r.Get("/suppliers/:id", middlewares.AuthMiddleware(), handlers.GetSupplierByID)
	r.Put("/suppliers/:id", middlewares.AuthMiddleware(), handlers.UpdateSupplier)
	r.Delete("/suppliers/:id", handlers.DeleteSupplier)
}
