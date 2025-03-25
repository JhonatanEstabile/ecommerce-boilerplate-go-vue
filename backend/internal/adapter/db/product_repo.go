package db

import (
	"database/sql"
	"ecommerce/internal/core/product"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) GetAll() ([]product.Product, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, stock_quantity, is_active FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []product.Product
	for rows.Next() {
		var p product.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.StockQuantity, &p.IsActive)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *PostgresRepo) Create(product product.Product) error {
	_, err := r.db.Exec(
		`
			INSERT INTO products
			("name", description, price, stock_quantity, is_active)
			VALUES($1, $2, $3, $4, $5)
		`,
		product.Name, product.Description, product.Price, product.StockQuantity, product.IsActive,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepo) GetOne(id int) (*product.Product, error) {
	var product product.Product

	err := r.db.QueryRow(`
		SELECT id, name, description, price, stock_quantity, is_active FROM products
		WHERE id = $1
	`, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.StockQuantity, &product.IsActive)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
