package service

import (
	"cleanarch/pkg/products/domain"
	"cleanarch/pkg/products/dto"
	"cleanarch/pkg/products/repository"
	"context"
	"errors"
)

type Service interface {
	CreateProducts(ctx context.Context, req dto.ProductsRequest) (*dto.Products, error)
	GetProductByid(ctx context.Context, id int) (*dto.Products, error)
	GetAllProducts(ctx context.Context) (*[]dto.Products, error)
	UpdateProductsById(ctx context.Context, id int, req dto.ProductsRequest) (int, error)
	DeleteProductsById(ctx context.Context, id int) (int, error)
	GetProductsItems(ctx context.Context) (*dto.ProductResponse, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateProducts(ctx context.Context, req dto.ProductsRequest) (*dto.Products, error) {
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
	product, err := s.repo.CreateProducts(ctx, productQuery)
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

func (s *service) GetAllProducts(ctx context.Context) (*[]dto.Products, error) {
	p, err := s.repo.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	var products []dto.Products
	for _, value := range *p {
		products = append(products, *value.ToModel())
	}
	return &products, nil
}

func (s *service) UpdateProductsById(ctx context.Context, id int, req dto.ProductsRequest) (int, error) {
	product := domain.ProductsQuery{
		Name:     req.Name,
		Price:    req.Price,
		Descript: req.Descript,
	}
	return s.repo.UpdateProductsById(ctx, id, product)
}

func (s *service) DeleteProductsById(ctx context.Context, id int) (int, error) {
	return s.repo.DeleteProductsById(ctx, id)
}

func (s *service) GetProductsItems(ctx context.Context) (*dto.ProductResponse, error) {
	p, err := s.repo.GetAllProducts(ctx)
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
