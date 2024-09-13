package transaction

import (
	infrafiber "nbid-online-shop/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	trxRoute := router.Group("transactions")
	{
		// menggunakan middleware
		trxRoute.Use(infrafiber.CheckAuth())

		// route akan berjalan  menggunakan middleware
		trxRoute.Post("/checkout", handler.CreateTransaction)
		trxRoute.Get("/user/histories", handler.GetTransactionsByUserPublicId)
	}
}
