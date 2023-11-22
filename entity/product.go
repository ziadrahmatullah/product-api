package entity

import "github.com/shopspring/decimal"

type Product struct {
	Id         int64           `json:"id"`
	Name       string          `json:"name"`
	Price      decimal.Decimal `json:"price"`
	CategoryId int64           `json:"category_id"`
	Stock      int             `json:"stock"`
}

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
