package kecamatan

import (
	"context"
	"nbid-online-shop/infra/response"
)

type Repository interface {
	GetVoterKecamatan(ctx context.Context, codeKecamatan string) (kecamatan Kecamatan, err error)
	GetAllVoter(ctx context.Context) (kecamatan Kecamatan, err error)
	GetListKecamatan(ctx context.Context) (kecamatans []Kecamatan, err error)
	GetListKelurahanFromKecamatan(ctx context.Context, codeKecamatan string) (kelurahans []Kelurahan, err error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) GetVoterKecamatan(ctx context.Context, codeKecamatan string) (kecamatan Kecamatan, err error) {

	model, err := s.repo.GetVoterKecamatan(ctx, codeKecamatan)
	if err != nil {
		return
	}

	return model, nil
}

func (s service) AllVoter(ctx context.Context) (kecamatan Kecamatan, err error) {

	model, err := s.repo.GetAllVoter(ctx)
	if err != nil {
		return
	}

	return model, nil
}

func (s service) GetListKecamatan(ctx context.Context) (kecamatans []Kecamatan, err error) {

	kecamatans, err = s.repo.GetListKecamatan(ctx)
	if err != nil {
		if err == response.ErrNotFound {
			return []Kecamatan{}, nil
		}
		return
	}

	if len(kecamatans) == 0 {
		return []Kecamatan{}, nil
	}
	return
}

func (s service) GetListKelurahanFromKecamatan(ctx context.Context, codeKecamatan string) (kelurahans []Kelurahan, err error) {

	kelurahans, err = s.repo.GetListKelurahanFromKecamatan(ctx, codeKecamatan)

	if err != nil {
		if err == response.ErrNotFound {
			return []Kelurahan{}, nil
		}
		return
	}

	if len(kelurahans) == 0 {
		return []Kelurahan{}, nil
	}
	return
}
