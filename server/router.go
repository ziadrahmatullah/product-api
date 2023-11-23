package server

import (
	// "net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/handler"
	"github.com/gin-gonic/gin"
	// "github.com/gorilla/mux"
)

type RouterOpts struct {
	ProductHandler *handler.ProductHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	product := router.Group("/products")
	product.GET("/", opts.ProductHandler.HandleListProducts)
	product.POST("/", opts.ProductHandler.HandleCreateProduct)
	product.PUT("/{id}", opts.ProductHandler.HandleUpdateProduct)
	product.DELETE("/{id}", opts.ProductHandler.HandleDeleteProduct)
	product.GET("/{id}", opts.ProductHandler.HandleGetProductById)
	
	// r := mux.NewRouter()
	// r.HandleFunc("/products", opts.ProductHandler.HandleListProducts).Methods(http.MethodGet)
	// r.HandleFunc("/products", opts.ProductHandler.HandleCreateProduct).Methods(http.MethodPost)
	// r.HandleFunc("/products/{id}", opts.ProductHandler.HandleUpdateProduct).Methods(http.MethodPut)
	// r.HandleFunc("/products/{id}", opts.ProductHandler.HandleDeleteProduct).Methods(http.MethodDelete)
	// r.HandleFunc("/products/{id}", opts.ProductHandler.HandleGetProductById).Methods(http.MethodGet)
	return router
}
