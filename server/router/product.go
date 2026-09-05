package router

import (
	"cleanarch/pkg/products/handler"
	"cleanarch/pkg/products/repository"
	"cleanarch/pkg/products/service"
	"database/sql"

	"github.com/gofiber/fiber/v3"
)

func Product(router fiber.Router, db *sql.DB) {
	productsRepo := repository.NewRepository(db)
	productsService := service.NewService(productsRepo)
	productsHandler := handler.NewHandler(productsService)

	group := router.Group("/product")
	group.Post("/", productsHandler.CreateProducts)
	group.Get("/", productsHandler.GetAllProducts)
	group.Get("/paginated", productsHandler.GetAllProductsPaginated)
	group.Get("/:id", productsHandler.GetProductByid)
	group.Patch("/:id", productsHandler.UpdateProductsById)
	group.Delete("/:id", productsHandler.DeleteProductsById)
	group.Get("/list", productsHandler.GetProductsItems)
}
