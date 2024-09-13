package kelurahan

import (
	infrafiber "nbid-online-shop/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	productRoute := router.Group("kelurahan")
	{

		productRoute.Get("/:code", infrafiber.CheckAuth(), handler.GetListTPSFromKelurahan)
		productRoute.Get("/voter/:code", infrafiber.CheckAuth(), handler.GetKelurahanData)
	}
}
