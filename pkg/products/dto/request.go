package dto

type ProductsRequest struct {
	Name     string `json:"name" validate:"required"`
	Price    int    `json:"price" validate:"required,gte=0"`
	Descript string `json:"descript"`
}
