package domain

import "cleanarch/pkg/products/dto"

type ProductsModel struct {
	Id       int
	Name     string
	Price    int
	Descript string
}

func (p *ProductsModel) ToModel() *dto.Products {
	return &dto.Products{
		Id:       p.Id,
		Name:     p.Name,
		Price:    p.Price,
		Descript: p.Descript,
	}
}
