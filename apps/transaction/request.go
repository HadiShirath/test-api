package transaction

type CreateTransactionRequestPayload struct {
	UserPublicId string `json:"-"`
	ProductSKU   string `json:"product_sku"`
	Amount       uint8  `json:"amount"`
}
