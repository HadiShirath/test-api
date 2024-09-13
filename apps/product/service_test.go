package product

import (
	"context"
	"log"
	"nbid-online-shop/external/database"
	"nbid-online-shop/infra/response"
	"nbid-online-shop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		return
	}

	repo := NewRepository(db)
	svc = NewService(repo)
}

func TestCreateProduct_Success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "Celana Jeans",
		Price: 12_000,
		Stock: 5,
	}

	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}

func TestCreateProduct_Failed(t *testing.T) {
	t.Run("Create Product Failed", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name:  "",
			Price: 12,
			Stock: 4,
		}

		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})

}

func TestListProduct_Success(t *testing.T) {
	pagination := ListProductsRequestPayload{
		Cursor: 0,
		Size:   10,
	}

	products, err := svc.ListProducts(context.Background(), pagination)
	require.Nil(t, err)
	require.NotNil(t, products)
	log.Printf("%+v", products)
}

func TestProductBySKU_Success(t *testing.T) {
	sku := "c2633922-bc12-4e26-9342-0275414ab508"
	product, err := svc.ProductDetail(context.Background(), sku)

	require.Nil(t, err)
	require.NotNil(t, product)
	log.Printf("%+v", product)
}
