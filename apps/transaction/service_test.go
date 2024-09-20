package transaction

import (
	"context"
	"nbid-online-shop/external/database"
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

func TestCreateTransaction(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := CreateTransactionRequestPayload{
			ProductSKU:   "e4b8c576-fbca-4075-a6d2-18d05680688b",
			Amount:       1,
			UserPublicId: "e6858fcf-0a78-414b-a773-2ce8c99eadb1",
		}

		err := svc.CreateTransaction(context.Background(), req)
		require.Nil(t, err)
	})

}

func TestTransactionByUserPublicId(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := CreateTransactionRequestPayload{
			UserPublicId: "e6858fcf-0a78-414b-a773-2ce8c99eadb1",
		}

		trxs, err := svc.TransactionHistories(context.Background(), req.UserPublicId)
		require.Nil(t, err)
		require.NotNil(t, trxs)

	})
}
