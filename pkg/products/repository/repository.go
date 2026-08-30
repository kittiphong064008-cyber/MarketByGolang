package repository

import (
	"cleanarch/pkg/products/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Repository interface {
	Create(ctx context.Context, p domain.ProductsQuery) (domain.ProductsModel, error)
	FindByID(ctx context.Context, id int) (*domain.ProductsModel, error)
	List(ctx context.Context) (*[]domain.ProductsModel, error)
	ListPaginated(ctx context.Context, page int, limit int) (*[]domain.ProductsModel, int, error)
	Update(ctx context.Context, id int, req domain.ProductsQuery) (int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository { //invert Dependency
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, p domain.ProductsQuery) (domain.ProductsModel, error) {
	query := "INSERT INTO products(name, price ,descript) VALUES ($1,$2,$3) RETURNING id" //ถูกแล้ว
	var id int
	err := r.db.QueryRowContext(ctx, query, p.Name, p.Price, p.Descript).Scan(&id)
	if err != nil {
		return domain.ProductsModel{}, fmt.Errorf("Failed to Create Product ")
	}
	return domain.ProductsModel{
		Id:       id,
		Name:     p.Name,
		Price:    p.Price,
		Descript: p.Descript,
	}, err
}

func (r *repository) FindByID(ctx context.Context, id int) (*domain.ProductsModel, error) {
	query := "SELECT id,name,price,descript FROM products WHERE id=$1"
	value := r.db.QueryRowContext(ctx, query, id)
	var p domain.ProductsModel
	err := value.Scan(&p.Id, &p.Name, &p.Price, &p.Descript)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Products Not Found Id = %d", id)
		}
		return nil, fmt.Errorf("Failed to Get Product By Id ")
	}
	return &p, nil
}

func (r *repository) List(ctx context.Context) (*[]domain.ProductsModel, error) {
	query := "SELECT id,name,price,descript FROM products ORDER BY id DESC"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to Get All Product ")
	}
	defer rows.Close()
	var product []domain.ProductsModel
	for rows.Next() {
		var p domain.ProductsModel
		err := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Descript)
		if err != nil {
			return nil, fmt.Errorf("Error to Scan Product Rows")
		}
		product = append(product, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error Iterating Product Rows")
	}
	return &product, nil
}

func (r *repository) ListPaginated(ctx context.Context, page int, limit int) (*[]domain.ProductsModel, int, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	var total int
	countQuery := "SELECT COUNT(*) FROM products"
	if err := r.db.QueryRowContext(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("Failed to count products")
	}

	offset := (page - 1) * limit
	query := "SELECT id,name,price,descript FROM products ORDER BY id DESC LIMIT $1 OFFSET $2"
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("Failed to Get All Product ")
	}
	defer rows.Close()

	var product []domain.ProductsModel
	for rows.Next() {
		var p domain.ProductsModel
		err := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Descript)
		if err != nil {
			return nil, 0, fmt.Errorf("Error to Scan Product Rows")
		}
		product = append(product, p)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("Error Iterating Product Rows")
	}
	return &product, total, nil
}

func (r *repository) Update(ctx context.Context, id int, req domain.ProductsQuery) (int, error) {
	query := "UPDATE products SET name=$1, price=$2, descript=$3 WHERE id = $4 "
	result, err := r.db.ExecContext(ctx, query, req.Name, req.Price, req.Descript, id)
	if err != nil {
		return 0, fmt.Errorf("Failed to Update Product Id = %d", id)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Failed to Get Rows Affected for Product Id = %d", id)
	}
	if rowsAffected == 0 {
		return 0, fmt.Errorf("No Product Found Id = %d", id)
	}
	return int(rowsAffected), nil
}

func (r *repository) Delete(ctx context.Context, id int) (int, error) {
	query := "DELETE FROM products WHERE id=$1"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, fmt.Errorf("Failed to Delete Product Id = %d", id)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Failed to Get Rows Affected for Product Id = %d", id)
	}
	if rowsAffected == 0 {
		return int(rowsAffected), fmt.Errorf("No Product Found Id = %d", id)
	}
	return int(rowsAffected), nil
}
