package usecase

import (
	"github.com/Backend_Shoppe/models"
	"github.com/Backend_Shoppe/repository"
)

type ProductUsecase interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetProducts() ([]models.Product, error)
	UpdateProduct(id string, product *models.Product) (*models.Product, error)
	DeleteProduct(id string) error
}

type productUsecase struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{productRepo: repo}
}

func (u *productUsecase) CreateProduct(product *models.Product) (*models.Product, error) {
	return u.productRepo.CreateProduct(product)
}

func (u *productUsecase) GetProducts() ([]models.Product, error) {
	return u.productRepo.GetProducts()
}

func (u *productUsecase) UpdateProduct(id string, product *models.Product) (*models.Product, error) {
	return u.productRepo.UpdateProduct(id, product)
}

func (u *productUsecase) DeleteProduct(id string) error {
	return u.productRepo.DeleteProduct(id)
}
