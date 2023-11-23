package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/repository"
)

type ProductUsecase interface {
	GetProducts() ([]entity.Product, error)
	CreateProduct(entity.Product) error
	UpdateProduct(entity.Product) error
	DeleteProduct(int64) error
	GetProductById(int64) (entity.Product, error)
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
	return p.productRepository.FindAllProducts()
}

func (p *productUsecase) CreateProduct(newProduct entity.Product) error {
	return p.productRepository.CreateNewProduct(newProduct)
}

func (p *productUsecase) UpdateProduct(updateProduct entity.Product) error{
	return p.productRepository.UpdateProduct(updateProduct)
}

func (p *productUsecase) DeleteProduct(id int64) error{
	return p.productRepository.DeleteProduct(id)
}

func (p *productUsecase) GetProductById(id int64) (entity.Product, error){
	return p.productRepository.FindProductById(id)
}

