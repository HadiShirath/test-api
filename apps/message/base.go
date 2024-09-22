package message

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

	productRoute := router.Group("message")
	{
		productRoute.Get("/inbox", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.GetInboxList)
		productRoute.Get("/outbox", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.GetOutboxList)
		productRoute.Post("/inbox", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.UploadInbox)
		productRoute.Post("/outbox", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.CreateMessage)
	}
}
