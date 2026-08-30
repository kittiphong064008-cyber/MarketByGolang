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
	GetAllCatagory(ctx context.Context) (*[]dto.CatagoryResponse, error)
	UpdateCatagoryById(ctx context.Context, id int, req dto.CatagoryUpdateRequest) (int, error)
	DeleteCatagoryById(ctx context.Context, id int) (int, error)
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

func (s *service) GetAllCatagory(ctx context.Context) (*[]dto.CatagoryResponse, error) {
	catagories, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	var response []dto.CatagoryResponse
	for _, value := range *catagories {
		response = append(response, *value.ToModel())
	}
	return &response, nil
}

func (s *service) UpdateCatagoryById(ctx context.Context, id int, req dto.CatagoryUpdateRequest) (int, error) {
	existingCatagory, err := s.repo.FindById(ctx, id)
	if err != nil {
		return 0, err
	}
	if req.Name != nil {
		existingCatagory.Name = *req.Name
	}
	catagory := domain.CatagoryQuery{
		Name: existingCatagory.Name,
	}
	return s.repo.Update(ctx, id, catagory)
}

func (s *service) DeleteCatagoryById(ctx context.Context, id int) (int, error) {
	return s.repo.Delete(ctx, id)
}
