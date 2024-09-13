package product

import (
	"nbid-online-shop/infra/response"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		product := Product{
			Name:  "Baju",
			Price: 12_000,
			Stock: 2,
		}

		err := product.Validate()
		require.Nil(t, err)
	})

	t.Run("product required", func(t *testing.T) {
		product := Product{
			Name:  "",
			Price: 12_000,
			Stock: 2,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})

	t.Run("product invalid", func(t *testing.T) {
		product := Product{
			Name:  "ha",
			Price: 12_000,
			Stock: 2,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductInvalid, err)
	})

	t.Run("price invalid", func(t *testing.T) {
		product := Product{
			Name:  "Baju Baru",
			Price: 0,
			Stock: 3,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)
	})

	t.Run("stock invalid", func(t *testing.T) {
		product := Product{
			Name:  "Baju baru",
			Price: 12_000,
			Stock: 0,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrStockInvalid, err)
	})
}
