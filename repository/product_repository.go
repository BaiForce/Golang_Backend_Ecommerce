package repository

import (
	"database/sql"
	"fmt"

	"github.com/Backend_Shoppe/models"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetProducts() ([]models.Product, error)
	UpdateProduct(id string, product *models.Product) (*models.Product, error)
	DeleteProduct(id string) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	query := "INSERT INTO products (name, price, description, category_id) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, product.Name, product.Price, product.Description, product.CategoryID)
	if err != nil {
		fmt.Printf("CreateProduct error: %v\n", err)
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("LastInsertId error: %v\n", err)
		return nil, err
	}
	product.ID = int(id)
	return product, nil
}


func (r *productRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	query := "SELECT id, name, price, description, category_id FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}


func (r *productRepository) UpdateProduct(id string, product *models.Product) (*models.Product, error) {
	query := "UPDATE products SET name = ?, price = ?, description = ?, category_id = ? WHERE id = ?"
	_, err := r.db.Exec(query, product.Name, product.Price, product.Description, product.CategoryID, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}


func (r *productRepository) DeleteProduct(id string) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
