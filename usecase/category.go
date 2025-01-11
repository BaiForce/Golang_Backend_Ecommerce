package usecase

import (
	"github.com/Backend_Shoppe/models"
	"github.com/Backend_Shoppe/repository"
)

type CategoryUsecase interface {
	CreateCategory(category *models.Category) (*models.Category, error)
	GetCategories() ([]models.Category, error)
	UpdateCategory(id string, category *models.Category) (*models.Category, error)
	DeleteCategory(id string) error
}

type categoryUsecase struct {
	categoryRepo repository.CategoryRepository
}

// Constructor function untuk usecase
func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{categoryRepo: repo}
}

// Implementasi CreateCategory
func (u *categoryUsecase) CreateCategory(category *models.Category) (*models.Category, error) {
	return u.categoryRepo.CreateCategory(category)
}

// Implementasi GetCategories
func (u *categoryUsecase) GetCategories() ([]models.Category, error) {
	return u.categoryRepo.GetCategories()
}

// Implementasi UpdateCategory
func (u *categoryUsecase) UpdateCategory(id string, category *models.Category) (*models.Category, error) {
	return u.categoryRepo.UpdateCategory(id, category)
}

// Implementasi DeleteCategory
func (u *categoryUsecase) DeleteCategory(id string) error {
	return u.categoryRepo.DeleteCategory(id)
}
