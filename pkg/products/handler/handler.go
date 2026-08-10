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
		return c.Status(400).JSON(fiber.Map{"Message": err.Error()})
	}
	_, err := h.service.CreateProducts(c.Context(), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Message": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"Message": "Create Products Success"})
}

func (h *handler) GetProductByid(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Message": err.Error()})
	}
	res, err := h.service.GetProductByid(c.Context(), id)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{"Message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"Message": "GET Success", "Value": res})
}

func (h *handler) GetAllProducts(c fiber.Ctx) error {
	res, err := h.service.GetAllProducts(c.Context())
	if err != nil {
		return c.Status(501).JSON(fiber.Map{"Message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"Message": "GET Success", "Value": res})
}

func (h *handler) UpdateProductsById(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Message": err.Error()})
	}
	var p dto.ProductsRequest
	if err := c.Bind().Body(&p); err != nil {
		return c.Status(500).JSON(fiber.Map{"Message": err.Error()})
	}
	rowsAffected, err := h.service.UpdateProductsById(c.Context(), id, p)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"Message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"Message": "Update Products Success", "rowsAffected": rowsAffected})
}

func (h *handler) DeleteProductsById(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Message": err.Error()})
	}
	rowsAffected, err := h.service.DeleteProductsById(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"Message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"Message": "Delete Products Sucess", "rowsAffected": rowsAffected})
}

func (h *handler) GetProductsItems(c fiber.Ctx) error {
	res, err := h.service.GetProductsItems(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"Message": "GET Success", "Value": res})
}
