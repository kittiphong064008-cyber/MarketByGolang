package router

import (
	"cleanarch/pkg/catagory/handler"
	"cleanarch/pkg/catagory/repository"
	"cleanarch/pkg/catagory/service"
	"database/sql"

	"github.com/gofiber/fiber/v3"
)

func Catagory(router fiber.Router, db *sql.DB) {
	catagoryRepository := repository.NewRepository(db)
	catagoryService := service.NewService(catagoryRepository)
	catagoryHandler := handler.NewHandler(catagoryService)

	group := router.Group("/catagory")
	group.Post("/", catagoryHandler.CreateCatagory)
	group.Get("/:id", catagoryHandler.GetCatagoryById)
}
