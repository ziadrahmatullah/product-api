package repository

import (
	"database/sql"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/entity"
)

type ProductRepository interface {
	FindAllProducts() ([]entity.Product, error)
	CreateNewProduct(entity.Product) error
	UpdateProduct(entity.Product) error
	DeleteProduct(int64) error
	FindProductById(int64) (entity.Product, error)
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
	sql := `
		SELECT product_id, product_name, quantity, price 
		FROM products`
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

func (p *productRepository) CreateNewProduct(newProduct entity.Product) (err error) {
	stmt, err := p.db.Prepare(`
		INSERT INTO products (product_category_id, product_name, quantity, price, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id`)
	if err != nil {
		return
	}
	_, err = stmt.Exec(newProduct.CategoryId, newProduct.Name, newProduct.Stock, newProduct.Price)
	if err != nil {
		return
	}
	return nil
}

func (p *productRepository) UpdateProduct(updateProduct entity.Product) (err error){
	stmt, err := p.db.Prepare(`
		UPDATE products 
		SET product_name = $1, quantity = $2, price = $3, product_category_id = $4 
		WHERE product_id = $5`)
	if err != nil {
		return
	}
	_, err = stmt.Exec(updateProduct.Name, updateProduct.Stock, updateProduct.Price, updateProduct.CategoryId, updateProduct.Id)
	if err != nil {
		return
	}
	return nil
}

func (p *productRepository) DeleteProduct(id int64) (err error) {
	sql := "DELETE FROM products WHERE product_id = $1"
	_, err = p.db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) FindProductById(id int64) (product entity.Product, err error){
	sql := `SELECT product_id, product_name, quantity, price FROM products p JOIN product_categories pc ON p.product_category_id = pc.product_category_id WHERE product_id = $1`
	row, err := p.db.Query(sql, id)
	if err != nil {
		return
	}
	defer row.Close()

	err = row.Scan(&product.Id, &product.Name, &product.Stock, &product.Price)
	if err != nil {
		return entity.Product{}, err
	}
	err = row.Err()
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}
