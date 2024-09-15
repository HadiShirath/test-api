package tps

import (
	"context"
	"log"
	"nbid-online-shop/apps/auth"
	"nbid-online-shop/infra/response"
	"nbid-online-shop/internal/config"
)

type Repository interface {
	CreatePhoto(ctx context.Context, model TPS, userId string) (err error)
	GetAddressTPSByUserId(ctx context.Context, userId string) (tps TPS, err error)
	GetAllVoterTPS(ctx context.Context) (tps TPS, err error)
	GetVoterTPS(ctx context.Context, codeTPS string) (tps TPS, err error)
	GetAllTPSSaksiWithPaginationCursor(ctx context.Context, model TPSPagination) (tpss []TPS, err error)
	EditTPSSaksi(ctx context.Context, model TPS, userId string) (err error)
	GetListTPS(ctx context.Context) (tpss []TPS, err error)
	EditVoteTPS(ctx context.Context, model TPS, tpsId string) (err error)
	// UpdateVoteTPSByUserId(ctx context.Context, model TPS, userId string) (err error)
	UpdateVoteTPSByUserId(ctx context.Context, model TPS, userId string) (updatedModel TPS, err error)
}

type service struct {
	repo     Repository
	repoAuth auth.Repository
}

func NewService(repo Repository, repoAuth auth.Repository) service {
	return service{
		repo:     repo,
		repoAuth: repoAuth,
	}
}

func (s service) CreatePhoto(ctx context.Context, userId string, filename string) (err error) {

	productEntity := NewFromCreatePhotoRequest(
		CreatePhotoRequestPayload{
			Photo: filename,
		})

	if err = s.repo.CreatePhoto(ctx, productEntity, userId); err != nil {
		return
	}

	return

}

func (s service) CheckPhoto(ctx context.Context, userId string) (photo string, err error) {
	model, err := s.repo.GetAddressTPSByUserId(ctx, userId)
	if err != nil {
		return
	}

	return model.Photo, nil
}

func (s service) TPSAdressDetail(ctx context.Context, userId string) (token string, err error) {

	model, err := s.repo.GetAddressTPSByUserId(ctx, userId)
	if err != nil {
		return
	}

	token, err = model.GenerateTokenData(config.Cfg.App.Encryption.JWTSecret)
	return
}

func (s service) GetAllVoterTPS(ctx context.Context) (tps TPS, err error) {
	model, err := s.repo.GetAllVoterTPS(ctx)

	if err != nil {
		return
	}

	return model, nil
}

func (s service) GetVoterTPS(ctx context.Context, codeTPS string) (tps TPS, err error) {

	model, err := s.repo.GetVoterTPS(ctx, codeTPS)
	if err != nil {
		return
	}

	return model, nil
}

func (s service) ListTPSSaksi(ctx context.Context, req ListTPSSaksiRequestPayload) (tpss []TPS, err error) {
	pagination := NewTPSSaksiPaginationFromProductRequest(req)

	tpss, err = s.repo.GetAllTPSSaksiWithPaginationCursor(ctx, pagination)
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

func (s service) EditTPSSaksi(ctx context.Context, req EditTPSSaksiRequestPayload, userId string) (err error) {
	saksiEntity := NewFromEditTPSSaksiRequest(req)

	if err = s.repo.EditTPSSaksi(ctx, saksiEntity, userId); err != nil {
		return
	}

	return

}

func (s service) GetListTPS(ctx context.Context) (tpss []TPS, err error) {

	tpss, err = s.repo.GetListTPS(ctx)

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

func (s service) EditVoteTPS(ctx context.Context, req EditVoteTPSRequestPayload, tpsId string) (err error) {
	voteTPSEntity := NewFromEditVoteTPSRequest(req)

	if err = s.repo.EditVoteTPS(ctx, voteTPSEntity, tpsId); err != nil {
		return
	}

	return
}

func (s service) UpdateVoteTPSByUserId(ctx context.Context, req EditVoteTPSBySaksiRequestPayload, username string) (tps TPS, err error) {

	voteTPSEntity := NewFromEditVoteBySaksiTPSRequest(req)

	if err = validateCodeUnique(req.CodeUnique); err != nil {
		return
	}

	log.Println(voteTPSEntity)

	model, err := s.repoAuth.GetAuthByUsername(ctx, username)
	if err != nil {
		return
	}

	tps, err = s.repo.UpdateVoteTPSByUserId(ctx, voteTPSEntity, model.PublicId.String())
	if err != nil {
		return
	}

	return
}
