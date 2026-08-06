package dto

type Products struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Descript string `json:"descript"`
}

type ProductResponse struct {
	TotalItem int
	MaxPrice  int
	MinPrice  int
	Products  []Products
}
