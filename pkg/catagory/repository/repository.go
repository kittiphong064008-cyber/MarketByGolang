package repository

import (
	"cleanarch/pkg/catagory/domain"
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	Create(ctx context.Context, c domain.CatagoryQuery) (domain.CatagoryModel, error)
	FindById(ctx context.Context, id int) (domain.CatagoryModel, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, c domain.CatagoryQuery) (domain.CatagoryModel, error) {
	query := "INSERT INTO catagory(name) VALUES ($1) RETURNING id"
	var id int
	err := r.db.QueryRowContext(ctx, query, c.Name).Scan(&id)
	if err != nil {
		return domain.CatagoryModel{}, fmt.Errorf("Failed to Create Catagory")
	}
	return domain.CatagoryModel{
		Id:   id,
		Name: c.Name,
	}, nil
}

func (r *repository) FindById(ctx context.Context, id int) (domain.CatagoryModel, error) {
	query := "SELECT id, name FROM catagory WHERE id = $1"
	var catagory domain.CatagoryModel
	err := r.db.QueryRowContext(ctx, query, id).Scan(&catagory.Id, &catagory.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.CatagoryModel{}, fmt.Errorf("Catagory not found")
		}
		return domain.CatagoryModel{}, fmt.Errorf("Failed to get Catagory by ID: %v", err)
	}
	return catagory, nil
}
