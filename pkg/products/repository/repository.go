package repository

import (
	"cleanarch/pkg/products/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Repository interface {
	CreateProducts(domain.ProductsQuery) (domain.ProductsModel, error)
	GetProductByid(ctx context.Context, id int) (*domain.ProductsModel, error)
	GetAllProducts() (*[]domain.ProductsModel, error)
	UpdateProductsById(ctx context.Context, id int, req domain.ProductsQuery) (int, error)
	DeleteProductsById(ctx context.Context, id int) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository { //invert Dependency
	return &repository{
		db: db,
	}
}

func (r *repository) CreateProducts(p domain.ProductsQuery) (domain.ProductsModel, error) {
	query := "INSERT INTO products(name, price ,descript) VALUES ($1,$2,$3)"
	result, err := r.db.Exec(query, p.Name, p.Price, p.Descript)
	id, err := result.LastInsertId()
	if err != nil {
		return domain.ProductsModel{}, err
	}
	return domain.ProductsModel{
		Id:       int(id),
		Name:     p.Name,
		Price:    p.Price,
		Descript: p.Descript,
	}, err
}

func (r *repository) GetProductByid(ctx context.Context, id int) (*domain.ProductsModel, error) {
	query := "SELECT id,name,price,descript FROM products WHERE id=$1"
	value := r.db.QueryRowContext(ctx, query, id)
	var p domain.ProductsModel
	err := value.Scan(&p.Id, &p.Name, &p.Price, &p.Descript) //โยน Adddress ให้ DATABASE แล้วใส่ค่า
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Products Not Found Id = %d", id)
		}
		return nil, err
	}
	return &p, nil
}

func (r *repository) GetAllProducts() (*[]domain.ProductsModel, error) {
	query := "SELECT id,name,price,descript FROM products ORDER BY id DESC"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var product []domain.ProductsModel
	for rows.Next() {
		var p domain.ProductsModel
		err := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Descript)
		if err != nil {
			return nil, err
		}
		product = append(product, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) UpdateProductsById(ctx context.Context, id int, req domain.ProductsQuery) (int, error) {
	query := "UPDATE products SET name=$1, price=$2, descript=$3 WHERE id = $4 "
	result, err := r.db.ExecContext(ctx, query, req.Name, req.Price, req.Descript, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rowsAffected == 0 {
		return 0, fmt.Errorf("No product found id = %d", id)
	}
	return int(rowsAffected), nil
}

func (r *repository) DeleteProductsById(ctx context.Context, id int) (int, error) {
	query := "DELETE FROM products WHERE id=$1"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rowsAffected == 0 {
		return int(rowsAffected), fmt.Errorf("No product found id = %d", id)
	}
	return int(rowsAffected), nil
}
