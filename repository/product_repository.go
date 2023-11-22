package repository

import (
	"database/sql"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/entity"
)

type ProductRepository interface {
	FindAllProducts() ([]entity.Product, error)
}

type productReposity struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productReposity{
		db: db,
	}
}

func (p *productReposity) FindAllProducts() ([]entity.Product, error) {
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
