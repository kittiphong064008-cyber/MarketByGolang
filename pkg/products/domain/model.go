package domain

type ProductsModel struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Descript string `json:"descript"`
}
