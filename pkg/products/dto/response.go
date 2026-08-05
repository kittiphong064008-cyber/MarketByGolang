package dto

type Products struct {
	Id       int
	Name     string
	Price    int
	Descript string
}

type ProductResponse struct {
	TotalItem int
	MaxPrice  int
	MinPrice  int
	Products  []Products
}
