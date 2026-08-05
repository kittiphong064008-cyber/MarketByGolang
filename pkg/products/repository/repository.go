package repository

import (
	"cleanarch/pkg/products/domain"
	"database/sql"
)

type Repository interface {
	CreateProducts(domain.ProductsQuery) error
	GetProductByid(id int) (*domain.ProductsModel, error)
	GetAllProducts() (*[]domain.ProductsModel, error)
	UpdateProductsById(id int, req domain.ProductsQuery) error
	DeleteProductsById(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateProducts(p domain.ProductsQuery) error {
	query := "INSERT INTO products(name, price ,descript) VALUES ($1,$2,$3)"
	_, err := r.db.Exec(query, p.Name, p.Price, p.Descript)
	return err
}

func (r *repository) GetProductByid(id int) (*domain.ProductsModel, error) {
	query := "SELECT id,name,price,descript FROM products WHERE id=$1"
	value := r.db.QueryRow(query, id)
	var p domain.ProductsModel
	err := value.Scan(&p.Id, &p.Name, &p.Price, &p.Descript)
	if err != nil {
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

func (r *repository) UpdateProductsById(id int, req domain.ProductsQuery) error {
	query := "UPDATE products SET name=$1, price=$2, descript=$3 WHERE id=$4"
	_, err := r.db.Exec(query, req.Name, req.Price, req.Descript, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteProductsById(id int) error {
	query := "DELETE FROM products WHERE id=$1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
