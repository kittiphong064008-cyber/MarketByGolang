package server

import (
	"cleanarch/pkg/products/handler"
	"cleanarch/pkg/products/repository"
	"cleanarch/pkg/products/service"
	"database/sql"

	"github.com/gofiber/fiber/v3"
)

func NewServer(db *sql.DB) *fiber.App {
	app := fiber.New()
	productsRepo := repository.NewRepository(db)
	productsService := service.NewService(productsRepo)
	productsHandler := handler.NewHandler(productsService)

	app.Post("/products", productsHandler.CreateProducts)
	app.Get("/products/:id", productsHandler.GetProductByid)
	app.Get("/products", productsHandler.GetAllProducts)
	app.Patch("/products/:id", productsHandler.UpdateProductsById)
	app.Delete("/products/:id", productsHandler.DeleteProductsById)

	app.Get("/products-items", productsHandler.GetProductsItems)

	return app
}
