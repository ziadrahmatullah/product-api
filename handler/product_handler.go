package handler

import (
	"encoding/json"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/usecase"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
	}
}

func (h *ProductHandler) HandleListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := h.productUsecase.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := dto.Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := dto.Response{
		Data: data,
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := dto.Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
}

func (h *ProductHandler) HandleCreateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	newProduct := entity.Product{}
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := dto.Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	err = h.productUsecase.CreateProduct(newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := dto.Response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := dto.Response{
		Data: newProduct,
	}
	json.NewEncoder(w).Encode(resp)
}
