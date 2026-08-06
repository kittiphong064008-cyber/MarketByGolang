package domain

type ProductsQuery struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Descript string `json:"descript"`
}
