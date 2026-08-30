package service

import (
	"cleanarch/pkg/products/domain"
	"cleanarch/pkg/products/dto"
	"cleanarch/pkg/products/repository"
	"context"
	//ย้ายไป handler
)

type Service interface {
	CreateProducts(ctx context.Context, req dto.ProductsRequest) (*dto.Products, error)
	GetProductByid(ctx context.Context, id int) (*dto.Products, error)
	GetAllProducts(ctx context.Context, page int, limit int) (*[]dto.Products, *dto.Pagination, error)
	UpdateProductsById(ctx context.Context, id int, req dto.ProductsUpdateRequest) (int, error)
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
	product, err := s.repo.Create(ctx, productQuery)
	if err != nil {
		return nil, err
	}
	return product.ToModel(), nil
}

func (s *service) GetProductByid(ctx context.Context, id int) (*dto.Products, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return p.ToModel(), nil
}

func (s *service) GetAllProducts(ctx context.Context, page int, limit int) (*[]dto.Products, *dto.Pagination, error) {
	p, total, err := s.repo.ListPaginated(ctx, page, limit)
	if err != nil {
		return nil, nil, err
	}
	var products []dto.Products
	for _, value := range *p {
		products = append(products, *value.ToModel())
	}
	pagination := dto.BuildPagination(page, limit, total)
	return &products, pagination, nil
}

func (s *service) UpdateProductsById(ctx context.Context, id int, req dto.ProductsUpdateRequest) (int, error) {
	existingProduct, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return 0, err
	}
	if req.Name != nil {
		existingProduct.Name = *req.Name
	}
	if req.Price != nil {
		existingProduct.Price = *req.Price
	}
	if req.Descript != nil {
		existingProduct.Descript = *req.Descript
	}
	product := domain.ProductsQuery{
		Name:     existingProduct.Name,
		Price:    existingProduct.Price,
		Descript: existingProduct.Descript,
	}
	return s.repo.Update(ctx, id, product)
}

func (s *service) DeleteProductsById(ctx context.Context, id int) (int, error) {
	return s.repo.Delete(ctx, id)
}

func (s *service) GetProductsItems(ctx context.Context) (*dto.ProductResponse, error) {
	p, err := s.repo.List(ctx)
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
