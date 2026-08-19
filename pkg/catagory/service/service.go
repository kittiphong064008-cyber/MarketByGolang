package service

import (
	"cleanarch/pkg/catagory/domain"
	"cleanarch/pkg/catagory/dto"
	"cleanarch/pkg/catagory/repository"
	"context"
)

type Service interface {
	CreateCatagory(ctx context.Context, req dto.CatagoryRequest) (*dto.CatagoryResponse, error)
	GetCatagoryById(ctx context.Context, id int) (*dto.CatagoryResponse, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateCatagory(ctx context.Context, req dto.CatagoryRequest) (*dto.CatagoryResponse, error) {
	catagoryRequest := domain.CatagoryQuery{
		Name: req.Name,
	}
	catagory, err := s.repo.Create(ctx, catagoryRequest)
	if err != nil {
		return nil, err
	}
	return catagory.ToModel(), nil
}

func (s *service) GetCatagoryById(ctx context.Context, id int) (*dto.CatagoryResponse, error) {
	catagory, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return catagory.ToModel(), nil
}
