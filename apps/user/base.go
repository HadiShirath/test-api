package user

import (
	"nbid-online-shop/apps/auth"
	infrafiber "nbid-online-shop/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	productRoute := router.Group("user")
	{
		productRoute.Put("/:id", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.EditAuth)
	}
}
