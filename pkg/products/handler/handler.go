package handler

import (
	"cleanarch/pkg/products/dto"
	"cleanarch/pkg/products/service"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateProducts(c fiber.Ctx) error {
	var req dto.ProductsRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"Message": "Invalid Value"})
	}
	err := h.service.CreateProducts(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Message": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"Message": "Creaate Products Success"})
}

func (h *handler) GetProductByid(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Message": err.Error()})
	}
	res, err := h.service.GetProductByid(id)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{"Message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"Message": "GET Success", "Value": res})
}

func (h *handler) GetAllProducts(c fiber.Ctx) error {
	res, err := h.service.GetAllProducts()
	if err != nil {
		return c.Status(501).JSON(fiber.Map{"Message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"Message": "GET Success", "Value": res})
}
