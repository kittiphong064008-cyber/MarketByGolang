package service

import (
	"cleanarch/pkg/products/domain"
	"cleanarch/pkg/products/dto"
	"cleanarch/pkg/products/repository"
	"context"
	"errors"
)

type Service interface {
	CreateProducts(req dto.ProductsRequest) (*dto.Products, error)
	GetProductByid(ctx context.Context, id int) (*dto.Products, error)
	GetAllProducts() (*[]dto.Products, error)
	UpdateProductsById(id int, req dto.ProductsRequest) error
	DeleteProductsById(id int) error
	GetProductsItems() (*dto.ProductResponse, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateProducts(req dto.ProductsRequest) (*dto.Products, error) {
	productQuery := domain.ProductsQuery{
		Name:     req.Name,
		Price:    req.Price,
		Descript: req.Descript,
	}
	if req.Price < 0 {
		return nil, errors.New("price is minus")
	}
	if req.Name == "" {
		return nil, errors.New("Name require")
	}
	product, err := s.repo.CreateProducts(productQuery)
	if err != nil {
		return nil, err
	}
	return product.ToModel(), nil
}

func (s *service) GetProductByid(ctx context.Context, id int) (*dto.Products, error) {
	p, err := s.repo.GetProductByid(ctx, id)
	if err != nil {
		return nil, err
	}

	return p.ToModel(), err
}

func (s *service) GetAllProducts() (*[]dto.Products, error) {
	p, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	var products []dto.Products
	for _, value := range *p {
		products = append(products, *value.ToModel())
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

func (s *service) GetProductsItems() (*dto.ProductResponse, error) {
	p, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	var product []dto.Products
	var maxPrice, minPrice int
	for _, value := range *p {
		product = append(product, *value.ToModel())

		if value.Price > maxPrice {
			maxPrice = value.Price
		}
		if value.Price < minPrice || minPrice == 0 {
			minPrice = value.Price
		}

	}
	products := dto.ProductResponse{
		TotalItem: len(product),
		MaxPrice:  maxPrice,
		MinPrice:  minPrice,
		Products:  product,
	}
	return &products, nil
}
