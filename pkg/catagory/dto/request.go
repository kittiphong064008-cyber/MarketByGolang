package dto

type CatagoryRequest struct {
	Name string `json:"name" validate:"required"`
}
