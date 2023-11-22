package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/repository"
)

type ProductUsecase interface {
	GetProducts() ([]entity.Product, error)
	CreateProduct(entity.Product) error
}

type productUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(p repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepository: p,
	}
}

func (p *productUsecase) GetProducts() ([]entity.Product, error) {
	products, err := p.productRepository.FindAllProducts()
	if err != nil {
		return []entity.Product{}, err
	}
	return products, nil
}

func (p *productUsecase) CreateProduct(newProduct entity.Product) error{
	return p.productRepository.CreateNewProducts(newProduct)
}
