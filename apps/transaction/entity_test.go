package transaction

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSetSubTotal(t *testing.T) {
	trx := Transaction{
		ProductPrice: 10_000,
		Amount:       2,
	}
	expected := uint(20_000)

	require.Equal(t, expected, trx.SubTotal)
}

func TestSetGrandTotal(t *testing.T) {
	t.Run("with platform fee", func(t *testing.T) {
		trx := Transaction{
			ProductPrice: 40_000,
			Amount:       3,
			PlatformFee:  2_000,
		}

		expected := uint(122_000)
		trx.SetGrandTotal()
		require.Equal(t, expected, trx.GrandTotal)
	})

	t.Run("without platform fee", func(t *testing.T) {
		trx := Transaction{
			ProductPrice: 20_000,
			Amount:       2,
		}

		expected := uint(40_000)

		trx.SetGrandTotal()
		require.Equal(t, expected, trx.GrandTotal)
	})
}

func TestSetProductJSON(t *testing.T) {
	products := Product{
		Id:    1,
		SKU:   uuid.NewString(),
		Name:  "Baju",
		Price: 12_000,
	}
	trx := Transaction{}
	err := trx.SetProductJSON(products)

	require.Nil(t, err)
	require.NotNil(t, trx)

	productFromTrx, err := trx.GetProduct()

	require.Nil(t, err)
	require.NotEmpty(t, productFromTrx)
	require.Equal(t, products, productFromTrx)
}

func TestTransactionStatus(t *testing.T) {
	type tableTest struct {
		title    string
		expected string
		trx      Transaction
	}

	var tableTests = []tableTest{
		{
			title:    "status created",
			trx:      Transaction{Status: TransactionStatus_Created},
			expected: TRX_CREATED,
		},
		{
			title:    "status on progress",
			trx:      Transaction{Status: TransactionStatus_Progress},
			expected: TRX_PROGRESS,
		},
		{
			title:    "status in delivery",
			trx:      Transaction{Status: TransactionStatus_InDelivery},
			expected: TRX_IN_DELIVERY,
		},
		{
			title:    "status completed",
			trx:      Transaction{Status: TransactionStatus_Complete},
			expected: TRX_COMPLETED,
		},
		{
			title:    "status unknown",
			trx:      Transaction{Status: TransactionStatus_Unknown},
			expected: TRX_UNKNOWN,
		},
	}

	for _, test := range tableTests {
		t.Run(test.title, func(t *testing.T) {
			require.Equal(t, test.expected, test.trx.GetStatus())

		})
	}
}
