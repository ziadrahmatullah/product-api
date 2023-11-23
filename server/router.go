package server

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/handler"
	"github.com/gorilla/mux"
)

type RouterOpts struct {
	ProductHandler *handler.ProductHandler
}

func NewRouter(opts RouterOpts) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/products", opts.ProductHandler.HandleListProducts).Methods(http.MethodGet)
	r.HandleFunc("/products", opts.ProductHandler.HandleCreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/products/{id}", opts.ProductHandler.HandleUpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/products/{id}", opts.ProductHandler.HandleDeleteProduct).Methods(http.MethodDelete)
	r.HandleFunc("/products/{id}", opts.ProductHandler.HandleGetProductById).Methods(http.MethodGet)
	return r
}
