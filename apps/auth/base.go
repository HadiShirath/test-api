package auth

import (
	infrafiber "nbid-online-shop/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	authRouter := router.Group("auth")
	{
		authRouter.Post("register", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(ROLE_Admin)}), handler.register)
		authRouter.Post("login", handler.login)
	}

}
