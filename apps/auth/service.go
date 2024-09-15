package auth

import (
	"context"
	"nbid-online-shop/infra/response"
	"nbid-online-shop/internal/config"
)

type Repository interface {
	GetAuthByUsername(ctx context.Context, email string) (model AuthEntity, err error)
	CreateAuth(ctx context.Context, model AuthEntity) (err error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) register(ctx context.Context, req RegisterRequestPayload) (err error) {
	authEntity := NewFromRegisterRequest(req)
	if err = authEntity.Validate(); err != nil {
		return
	}

	if err = authEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
		return
	}

	model, _ := s.repo.GetAuthByUsername(ctx, authEntity.Username)

	if model.IsExists() {
		return response.ErrUsernameAlreadyUsed
	}

	return s.repo.CreateAuth(ctx, authEntity)
}

func (s service) login(ctx context.Context, req LoginRequestPayload) (token string, role string, err error) {
	authEntity := NewFromLoginRequest(req)

	if err = authEntity.ValidateEmail(); err != nil {
		return
	}
	if err = authEntity.ValidatePassword(); err != nil {
		return
	}

	model, err := s.repo.GetAuthByUsername(ctx, authEntity.Username)
	if err != nil {
		if err != response.ErrNotFound {
			return
		}
		return
	}

	if err = authEntity.VerifyPasswordFromPlain(model.Password); err != nil {
		err = response.ErrPasswordNotMatch
		return
	}

	token, err = model.GenerateToken(config.Cfg.App.Encryption.JWTSecret)
	return token, string(model.Role), err

}
