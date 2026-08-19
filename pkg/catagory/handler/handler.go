package handler

import (
	"cleanarch/pkg/catagory/dto"
	"cleanarch/pkg/catagory/service"
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

func (h *handler) CreateCatagory(c fiber.Ctx) error {
	var req dto.CatagoryRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: err.Error()})
	}
	if err := val.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: "Invalid Body Format"})
	}
	res, err := h.service.CreateCatagory(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{Message: err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ResponseOne{Message: "Create Success", Value: *res})
}

func (h *handler) GetCatagoryById(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseError{Message: "Invalid ID"})
	}
	res, err := h.service.GetCatagoryById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseError{Message: err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(dto.ResponseOne{Message: "Get Success", Value: *res})
}
