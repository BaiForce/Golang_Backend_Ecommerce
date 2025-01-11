package repository

import (
	"database/sql"
	"fmt"

	"github.com/Backend_Shoppe/models"
)

type CategoryRepository interface {
	CreateCategory(category *models.Category) (*models.Category, error)
	GetCategories() ([]models.Category, error)
	UpdateCategory(id string, category *models.Category) (*models.Category, error)
	DeleteCategory(id string) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) CreateCategory(category *models.Category) (*models.Category, error) {
	query := "INSERT INTO categories (name, description) VALUES (?, ?)"
	result, err := r.db.Exec(query, category.Name, category.Description)
	if err != nil {
		fmt.Printf("CreateCategory error: %v\n", err)
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("LastInsertId error: %v\n", err)
		return nil, err
	}
	category.ID = int(id)
	return category, nil
}

func (r *categoryRepository) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	query := "SELECT id, name, description FROM categories"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *categoryRepository) UpdateCategory(id string, category *models.Category) (*models.Category, error) {
	query := "UPDATE categories SET name = ?, description = ? WHERE id = ?"
	_, err := r.db.Exec(query, category.Name, category.Description, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) DeleteCategory(id string) error {
	query := "DELETE FROM categories WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
