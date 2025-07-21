package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestProductsDefault_Get(t *testing.T) {
	t.Run("should return products successfully", func(t *testing.T) {

		ctrl := gomock.NewController(t)

		mockRepo := mocks.NewMockRepositoryProducts(ctrl)

		expectedProducts := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Test Product", Price: 99.99, SellerId: 123}},
		}

		mockRepo.EXPECT().
			SearchProducts(gomock.Any()).
			Return(expectedProducts, nil).
			Times(1)

		h := handler.NewProductsDefault(mockRepo)
		req := httptest.NewRequest(http.MethodGet, "/product?id=1", nil)
		rr := httptest.NewRecorder()

		handlerFunc := h.Get()
		handlerFunc(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
		expectedBody := `{"data":{"1":{"id":1,"description":"Test Product","price":99.99,"seller_id":123}},"message":"success"}`
		require.JSONEq(t, expectedBody, rr.Body.String())
	})

	t.Run("should return 500 on repository error using Uber's mockgen", func(t *testing.T) {

		ctrl := gomock.NewController(t)

		mockRepo := mocks.NewMockRepositoryProducts(ctrl)

		mockRepo.EXPECT().
			SearchProducts(gomock.Any()).
			Return(nil, errors.New("database error")).
			Times(1)

		h := handler.NewProductsDefault(mockRepo)
		req := httptest.NewRequest(http.MethodGet, "/product", nil)
		rr := httptest.NewRecorder()

		handlerFunc := h.Get()
		handlerFunc(rr, req)

		require.Equal(t, http.StatusInternalServerError, rr.Code)
	})
}
