package repository

import (
	"cleanarch/pkg/catagory/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Repository interface {
	Create(ctx context.Context, c domain.CatagoryQuery) (domain.CatagoryModel, error)
	FindById(ctx context.Context, id int) (domain.CatagoryModel, error)
	List(ctx context.Context) (*[]domain.CatagoryModel, error)
	Update(ctx context.Context, id int, req domain.CatagoryQuery) (int, error)
	Delete(ctx context.Context, id int) (int, error)
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
		if errors.Is(err, sql.ErrNoRows) {
			return domain.CatagoryModel{}, fmt.Errorf("Catagory not found")
		}
		return domain.CatagoryModel{}, fmt.Errorf("Failed to get Catagory by ID: %v", err)
	}
	return catagory, nil
}

func (r *repository) List(ctx context.Context) (*[]domain.CatagoryModel, error) {
	query := "SELECT id,name FROM catagory ORDER BY id DESC"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to Get All Catagory")
	}
	defer rows.Close()

	var catagories []domain.CatagoryModel
	for rows.Next() {
		var catagory domain.CatagoryModel
		err := rows.Scan(&catagory.Id, &catagory.Name)
		if err != nil {
			return nil, fmt.Errorf("Error to Scan Catagory Rows")
		}
		catagories = append(catagories, catagory)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error Iterating Catagory Rows")
	}
	return &catagories, nil
}

func (r *repository) Update(ctx context.Context, id int, req domain.CatagoryQuery) (int, error) {
	query := "UPDATE catagory SET name=$1 WHERE id = $2"
	result, err := r.db.ExecContext(ctx, query, req.Name, id)
	if err != nil {
		return 0, fmt.Errorf("Failed to Update Catagory Id = %d", id)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Failed to Get Rows Affected for Catagory Id = %d", id)
	}
	if rowsAffected == 0 {
		return 0, fmt.Errorf("No Catagory Found Id = %d", id)
	}
	return int(rowsAffected), nil
}

func (r *repository) Delete(ctx context.Context, id int) (int, error) {
	query := "DELETE FROM catagory WHERE id=$1"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, fmt.Errorf("Failed to Delete Catagory Id = %d", id)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Failed to Get Rows Affected for Catagory Id = %d", id)
	}
	if rowsAffected == 0 {
		return int(rowsAffected), fmt.Errorf("No Catagory Found Id = %d", id)
	}
	return int(rowsAffected), nil
}
