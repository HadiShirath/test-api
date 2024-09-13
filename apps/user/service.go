package user

import (
	"context"
	"nbid-online-shop/internal/config"
)

type Repository interface {
	EditTPSSaksi(ctx context.Context, model User, userId string) (err error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) EditTPSSaksi(ctx context.Context, req EditUserRequestPayload, userId string) (err error) {
	saksiEntity := NewFromEditTPSSaksiRequest(req)

	if saksiEntity.Password != "" {
		if err = saksiEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
			return
		}
	}

	if err = s.repo.EditTPSSaksi(ctx, saksiEntity, userId); err != nil {
		return
	}

	return

}