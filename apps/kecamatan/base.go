package kecamatan

import (
	infrafiber "nbid-online-shop/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	productRoute := router.Group("kecamatan")
	{
		productRoute.Get("/", infrafiber.CheckAuth(), handler.GetAllVoter)
		productRoute.Get("/all", infrafiber.CheckAuth(), handler.ListKecamatan)
		productRoute.Get("/:code", infrafiber.CheckAuth(), handler.GetListKelurahanFromKecamatan)
		productRoute.Get("/voter/:code", infrafiber.CheckAuth(), handler.GetVoterKecamatan)
	}
}
