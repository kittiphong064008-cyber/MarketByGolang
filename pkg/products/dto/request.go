package dto

type ProductsRequest struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Descript string `json:"descript"`
}
