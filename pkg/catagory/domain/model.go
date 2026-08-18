package domain

import "cleanarch/pkg/catagory/dto"

type CatagoryModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (p *CatagoryModel) ToModel() *dto.CatagoryResponse {
	return &dto.CatagoryResponse{
		Id:   p.Id,
		Name: p.Name,
	}
}
