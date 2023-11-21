package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/database"
	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
)

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

type Response struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

var db *sql.DB

/*
TODO:
Get  all products
post a product
get product detail by id (join with category)
update a product by id
delete a product by id

Handle primary key
handle more error
refactor package path
*/

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := ListProducts(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := Response{
		Data: data,
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newProduct := Product{}
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	err = CreateProduct(db, newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := Response{
		Data: newProduct,
	}
	json.NewEncoder(w).Encode(resp)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	updateProduct := Product{}
	err := json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	err = UpdateProduct(db, updateProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := Response{
		Data: updateProduct,
	}
	json.NewEncoder(w).Encode(resp)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}

func ListProducts(db *sql.DB) ([]Product, error) {
	res := []Product{}
	sql := `SELECT product_id, product_name, quantity, price FROM products`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product Product

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

func CreateProduct(db *sql.DB, newProduct Product) error {
	stmt, err := db.Prepare("INSERT INTO products (product_category_id, product_name, quantity, price, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(newProduct.CategoryId, newProduct.Name, newProduct.Stock, newProduct.Price)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *sql.DB, updateProduct Product) error {
	stmt, err := db.Prepare("UPDATE products SET product_name = $1, quantity = $2, price = $3, product_category_id = $4 WHERE product_id = $5") 
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(updateProduct.Name, updateProduct.Stock, updateProduct.Price, updateProduct.CategoryId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(db *sql.DB, id int) error {

	sql := "DELETE FROM products WHERE id = $1"

	_, err := db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	db = database.InitDB()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/products", GetProductHandler).Methods("GET")
	r.HandleFunc("/products", CreateProductHandler).Methods("POST")

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	srv.ListenAndServe()
}
