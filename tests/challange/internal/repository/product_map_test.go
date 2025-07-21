package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductsMap_SearchProducts(t *testing.T) {
	t.Run("should return all products when query is empty", func(t *testing.T) {

		db := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 10.0, SellerId: 100}},
			2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Product 2", Price: 20.0, SellerId: 101}},
		}
		repo := repository.NewProductsMap(db)
		query := internal.ProductQuery{}

		result, err := repo.SearchProducts(query)

		require.NoError(t, err)
		require.Len(t, result, 2)
	})

	t.Run("should return product by id when query has id", func(t *testing.T) {

		db := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 10.0, SellerId: 100}},
			2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Product 2", Price: 20.0, SellerId: 101}},
		}
		repo := repository.NewProductsMap(db)
		query := internal.ProductQuery{Id: 1}

		result, err := repo.SearchProducts(query)

		require.NoError(t, err)
		require.Len(t, result, 1)
		require.Equal(t, "Product 1", result[1].Description)
	})

	t.Run("should return empty map when no product matches the id", func(t *testing.T) {

		db := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 10.0, SellerId: 100}},
		}
		repo := repository.NewProductsMap(db)
		query := internal.ProductQuery{Id: 99}

		result, err := repo.SearchProducts(query)

		require.NoError(t, err)
		require.Empty(t, result)
	})
}
