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

	kecamatanRoute := router.Group("kecamatan")
	{
		kecamatanRoute.Get("/", infrafiber.CheckAuth(), handler.GetAllVoter)
		kecamatanRoute.Get("/all", infrafiber.CheckAuth(), handler.ListKecamatan)
		kecamatanRoute.Get("/:code", infrafiber.CheckAuth(), handler.GetListKelurahanFromKecamatan)
		kecamatanRoute.Get("/voter/:code", infrafiber.CheckAuth(), handler.GetVoterKecamatan)
	}
}
