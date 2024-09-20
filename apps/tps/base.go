package tps

import (
	"nbid-online-shop/apps/auth"
	infrafiber "nbid-online-shop/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := NewRepository(db)
	repoAuth := auth.NewRepository(db)
	svc := NewService(repo, repoAuth)
	handler := NewHandler(svc)

	productRoute := router.Group("tps")
	{
		productRoute.Get("/", infrafiber.CheckAuth(), handler.TPSAdressDetail)
		productRoute.Get("/all", infrafiber.CheckAuth(), handler.GetListTPS)
		productRoute.Get("/saksi", infrafiber.CheckAuth(), handler.GetTPSSaksiPagination)
		productRoute.Post("/photo", infrafiber.CheckAuth(), handler.CreatePhoto)
		productRoute.Post("/upload", infrafiber.CheckAuth(), handler.UploadDataTPS)
		productRoute.Get("/voter/all", infrafiber.CheckAuth(), handler.GetAllVoterTPS)
		productRoute.Get("/voter/:code", infrafiber.CheckAuth(), handler.GetVoterTPS)
		productRoute.Put("/voter/user/:id", infrafiber.CheckAuth(), handler.UpdateVoteTPSByUser)
		productRoute.Put("/voter/:id", infrafiber.CheckAuth(), handler.UpdateVoteTPS)
	}
}
