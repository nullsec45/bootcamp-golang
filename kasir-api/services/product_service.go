package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
	"errors"
	"fmt"
)

type ProductService struct {
	repo *repositories.ProductRepository
	categoryRepo *repositories.CategoryRepository
}

func NewProductService(repo *repositories.ProductRepository, categoryRepo *repositories.CategoryRepository) *ProductService {
	return &ProductService{repo:repo, categoryRepo: categoryRepo}
}

func (s *ProductService) GetAll(name string) ([]models.Product, error) {
	return s.repo.GetAll(name)
}

func (s *ProductService) Create(data *models.Product) error {
	cat, err := s.categoryRepo.GetByID(data.Category.ID)
    if err != nil {
		fmt.Println(err)
        return errors.New("gagal membuat produk: kategori tidak ditemukan")
    }

	data.Category = *cat

	return s.repo.Create(data)
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(product *models.Product) error {
	cat, err := s.categoryRepo.GetByID(product.Category.ID)
    if err != nil {
        return errors.New("gagal update produk: kategori tidak ditemukan")
    }

	product.Category = *cat


	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}