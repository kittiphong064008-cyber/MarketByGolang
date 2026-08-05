package server

import (
	configs "cleanarch/config"
	"cleanarch/pkg/products/handler"
	"cleanarch/pkg/products/repository"
	"cleanarch/pkg/products/service"
	"database/sql"

	"github.com/gofiber/fiber/v3"
)

func NewServer(db *sql.DB, cfg configs.Configs) *fiber.App {
	app := fiber.New()
	productsRepo := repository.NewRepository(db)
	productsService := service.NewService(productsRepo)
	productsHandler := handler.NewHandler(productsService)

	app.Post("/products", productsHandler.CreateProducts)
	app.Get("/products/:id", productsHandler.GetProductByid)
	app.Get("/products", productsHandler.GetAllProducts)
	app.Patch("/products/:id", productsHandler.UpdateProductsById)
	return app
}
