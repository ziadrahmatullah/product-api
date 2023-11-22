package main

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/database"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/usecase"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	pr := repository.NewProductRepository(db)
	pu := usecase.NewProductUsecase(pr)
	ph := handler.NewProductHandler(pu)

	opts := server.RouterOpts{
		ProductHandler: ph,	
	}
	r := server.NewRouter(opts)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	srv.ListenAndServe()
}
