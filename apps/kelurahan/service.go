package kelurahan

import (
	"context"
	"nbid-online-shop/infra/response"
)

type Repository interface {
	GetKelurahanData(ctx context.Context, codeKelurahan string) (kelurahan Kelurahan, err error)
	GetListTPSFromKelurahan(ctx context.Context, codeKelurahan string) (tpss []TPS, err error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) GetKeluharanData(ctx context.Context, codeKelurahan string) (kelurahan Kelurahan, err error) {

	model, err := s.repo.GetKelurahanData(ctx, codeKelurahan)
	if err != nil {
		return
	}

	return model, nil
}

func (s service) GetListTPSFromKelurahan(ctx context.Context, codeKelurahan string) (tpss []TPS, err error) {

	tpss, err = s.repo.GetListTPSFromKelurahan(ctx, codeKelurahan)

	if err != nil {
		if err == response.ErrNotFound {
			return []TPS{}, nil
		}
		return
	}

	if len(tpss) == 0 {
		return []TPS{}, nil
	}
	return
}
