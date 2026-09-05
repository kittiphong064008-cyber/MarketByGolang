package handler

import (
	"cleanarch/pkg/products/dto"
	"cleanarch/pkg/products/service"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type handler struct {
	service service.Service
}

var val = validator.New()

func NewHandler(service service.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateProducts(c fiber.Ctx) error {
	var req dto.ProductsRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: err.Error()})
	}
	if err := val.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: "Invalid Body Format"})
	}
	res, err := h.service.CreateProducts(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{Message: err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ResponseOne{Message: "Create Success", Value: *res})
}

func (h *handler) GetProductByid(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: err.Error()})
	}
	res, err := h.service.GetProductByid(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ResponseError{Message: err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(dto.ResponseOne{Message: "GET Success", Value: *res})
}

func (h *handler) GetAllProducts(c fiber.Ctx) error {
	res, err := h.service.GetAllProducts(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{Message: err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(dto.ResponseAll{Message: "GET Success", Value: *res})
}

func (h *handler) GetAllProductsPaginated(c fiber.Ctx) error {
	var req dto.ProductsPaginationRequest
	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: "Invalid query format"})
	}
	if err := val.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: "page and limit must be greater than 0"})
	}
	res, pagination, err := h.service.GetAllProductsPaginated(c.Context(), req.Page, req.Limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{Message: err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(dto.ResponsePaginated{Message: "GET Success", Value: *res, Pagination: *pagination})
}

func (h *handler) UpdateProductsById(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: err.Error()})
	}
	var req dto.ProductsUpdateRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: err.Error()})
	}
	if err := val.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: "Invalid Body Format"})
	}
	rowsAffected, err := h.service.UpdateProductsById(c.Context(), id, req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ResponseError{Message: err.Error()})
	}
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(dto.ResponseError{Message: "Product not found"})
	}
	return c.Status(fiber.StatusOK).JSON(dto.ResponseRows{Message: "Update Success", RowsAffect: rowsAffected})
}

func (h *handler) DeleteProductsById(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: err.Error()})
	}
	rowsAffected, err := h.service.DeleteProductsById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ResponseError{Message: err.Error()})
	}
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(dto.ResponseError{Message: "Product not found"})
	}
	return c.Status(fiber.StatusOK).JSON(dto.ResponseRows{Message: "Delete Success", RowsAffect: rowsAffected})
}

func (h *handler) GetProductsItems(c fiber.Ctx) error {
	res, err := h.service.GetProductsItems(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{Message: err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(dto.ResponseProducts{Message: "Get Success", Value: *res})
}
