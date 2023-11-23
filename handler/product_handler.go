package handler

import (
	// "encoding/json"
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/usecase"
	"github.com/gin-gonic/gin"
	// "github.com/gorilla/mux"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
	}
}

func (h *ProductHandler) HandleListProducts(ctx *gin.Context) {
	resp := dto.Response{}
	products, err := h.productUsecase.GetProducts()
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": products})

	// w.Header().Set("Content-Type", "application/json")
	// jsonEncoder := json.NewEncoder(w)
	// data, err := h.productUsecase.GetProducts()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	resp := dto.Response{
	// 		Message: err.Error(),
	// 	}
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }
	// resp := dto.Response{
	// 	Data: data,
	// }
	// err = jsonEncoder.Encode(resp)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	resp := dto.Response{
	// 		Message: err.Error(),
	// 	}
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }
}

func (h *ProductHandler) HandleCreateProduct(ctx *gin.Context) {
	resp := dto.Response{}
	newProduct := entity.Product{}
	err := ctx.ShouldBindJSON(&newProduct)
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	err = h.productUsecase.CreateProduct(newProduct)
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Data = newProduct
	ctx.JSON(http.StatusOK, resp)

	// w.Header().Set("Content-Type", "application/json")
	// newProduct := entity.Product{}
	// err := json.NewDecoder(r.Body).Decode(&newProduct)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	resp := dto.Response{
	// 		Message: err.Error(),
	// 	}
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }
	// err = h.productUsecase.CreateProduct(newProduct)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	resp := dto.Response{
	// 		Message: err.Error(),
	// 	}
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }
	// resp := dto.Response{
	// 	Data: newProduct,
	// }
	// json.NewEncoder(w).Encode(resp)
}

func (h *ProductHandler) HandleUpdateProduct(ctx *gin.Context) {
	resp := dto.Response{}
	updateProduct := entity.Product{}
	err := ctx.ShouldBindJSON(&updateProduct)
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	err = h.productUsecase.UpdateProduct(updateProduct)
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Data = updateProduct
	ctx.JSON(http.StatusOK, resp)

	// w.Header().Set("Content-Type", "application/json")
	// updateProduct := entity.Product{}
	// err := json.NewDecoder(r.Body).Decode(&updateProduct)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	resp := dto.Response{
	// 		Message: err.Error(),
	// 	}
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }
	// err = h.productUsecase.UpdateProduct(updateProduct)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	resp := dto.Response{
	// 		Message: err.Error(),
	// 	}
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }
	// resp := dto.Response{
	// 	Data: updateProduct,
	// }
	// json.NewEncoder(w).Encode(resp)
}

func (h *ProductHandler) HandleDeleteProduct(ctx *gin.Context) {
	resp := dto.Response{}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	err = h.productUsecase.DeleteProduct(int64(id))
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Data = id
	ctx.JSON(http.StatusOK, resp)

	// w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// id, err := strconv.Atoi(params["id"])
	// if err != nil {
	// 	resp := dto.Response{
	// 		Message: err.Error(),
	// 	}
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }
	// err = h.productUsecase.DeleteProduct(int64(id))
	// if err != nil {
	// 	resp := dto.Response{
	// 		Message: err.Error(),
	// 	}
	// 	w.WriteHeader(http.StatusNotFound)
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }
	// resp := dto.Response{
	// 	Message: "Success deleted",
	// }
	// json.NewEncoder(w).Encode(resp)
}

func (h *ProductHandler) HandleGetProductById(ctx *gin.Context) {
	resp := dto.Response{}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	product, err := h.productUsecase.GetProductById(int64(id))
	if err != nil {
		resp.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Data = product
	ctx.JSON(http.StatusOK, resp)
	// w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// id, err := strconv.Atoi(params["id"])
	// if err != nil {
	// 	resp := dto.Response{
	// 		Message: err.Error(),
	// 	}
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }
	// data, err := h.productUsecase.GetProductById(int64(id))
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	resp := dto.Response{
	// 		Message: err.Error(),
	// 	}
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }
	// resp := dto.Response{
	// 	Data: data,
	// }
	// err = json.NewEncoder(w).Encode(resp)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	resp := dto.Response{
	// 		Message: err.Error(),
	// 	}
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }
}
