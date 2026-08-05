package service

import (
	"cleanarch/pkg/products/domain"
	"cleanarch/pkg/products/dto"
	"cleanarch/pkg/products/repository"
)

type Service interface {
	CreateProducts(dto.ProductsRequest) error
	GetProductByid(id int) (*dto.Products, error)
	GetAllProducts() (*[]dto.Products, error)
	UpdateProductsById(id int, req dto.ProductsRequest) error
	DeleteProductsById(id int) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateProducts(req dto.ProductsRequest) error {
	product := domain.ProductsQuery{
		Name:     req.Name,
		Price:    req.Price,
		Descript: req.Descript,
	}
	return s.repo.CreateProducts(product)
}

func (s *service) GetProductByid(id int) (*dto.Products, error) {
	p, err := s.repo.GetProductByid(id)
	if err != nil {
		return nil, err
	}
	products := dto.Products{
		Id:       p.Id,
		Name:     p.Name,
		Price:    p.Price,
		Descript: p.Descript,
	}
	return &products, err
}

func (s *service) GetAllProducts() (*[]dto.Products, error) {
	p, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	var products []dto.Products
	for _, value := range *p {
		product := dto.Products{
			Id:       value.Id,
			Name:     value.Name,
			Price:    value.Price,
			Descript: value.Descript,
		}
		products = append(products, product)
	}
	return &products, nil
}

func (s *service) UpdateProductsById(id int, req dto.ProductsRequest) error {
	product := domain.ProductsQuery{
		Name:     req.Name,
		Price:    req.Price,
		Descript: req.Descript,
	}
	return s.repo.UpdateProductsById(id, product)
}

func (s *service) DeleteProductsById(id int) error {
	return s.repo.DeleteProductsById(id)
}
