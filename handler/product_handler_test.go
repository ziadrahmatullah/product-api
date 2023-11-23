package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/product-api/-/tree/ziad-rahmatullah/server"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var products = []entity.Product{
	{
		Id:         1,
		Name:       "Sabun",
		Price:      decimal.NewFromInt(3500),
		CategoryId: 1,
		Stock:      999,
	},
}

func removeNewLine(str string) string {
	return strings.Trim(str, "\n")
}

// type ProductHandlerTestSuite struct {
// 	suite.Suite
// 	pu     *mocks.ProductUsecase
// 	ph     *handler.ProductHandler
// 	router http.Handler
// }

// func (suite *ProductHandlerTestSuite) SetupSubTest() {
// 	suite.pu = mocks.NewProductUsecase(suite.T())
// 	suite.ph = handler.NewProductHandler(suite.pu, validator.New())
// 	opts := server.RouterOpts{
// 		ProductHandler: suite.ph,
// 	}
// 	suite.router = server.NewRouter(opts)
// }



func TestProductHandler_HandleListProducts(t *testing.T) {
	t.Run("should return 200 with product list", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: products,
		})
		pu := mocks.NewProductUsecase(t)
		ph := handler.NewProductHandler(pu)
		pu.On("GetProducts").Return(products, nil)
		opts := server.RouterOpts{
			ProductHandler: ph,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/products", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), removeNewLine(rec.Body.String()))
	})

}
