package repository

import (
	"database/sql"
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/entity"
)

type ProductRepository interface {
	FindAllProducts() ([]entity.Product, error)
	CreateNewProducts(entity.Product) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (p *productRepository) FindAllProducts() ([]entity.Product, error) {
	res := []entity.Product{}
	sql := `SELECT product_id, product_name, quantity, price FROM products`
	rows, err := p.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Stock, &product.Price)
		if err != nil {
			return nil, err
		}
		res = append(res, product)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *productRepository) CreateNewProducts(newProduct entity.Product) error{
	stmt, err := p.db.Prepare("INSERT INTO products (product_category_id, product_name, quantity, price, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(newProduct.CategoryId, newProduct.Name, newProduct.Stock, newProduct.Price)
	if err != nil {
		return err
	}
	return nil
}