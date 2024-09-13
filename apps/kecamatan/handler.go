package kecamatan

import (
	"context"
	infrafiber "nbid-online-shop/infra/fiber"
	"nbid-online-shop/infra/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func NewHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) GetAllVoter(ctx *fiber.Ctx) error {

	kecamatan, err := h.svc.AllVoter(context.Background())
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	getAllVoterResponse := kecamatan.ToAllVoterResponse()

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get voter kelurahan success"),
		infrafiber.WithPayload(getAllVoterResponse),
	).Send(ctx)

}

func (h handler) GetVoterKecamatan(ctx *fiber.Ctx) error {
	codeKecamatan := ctx.Params("code", "")
	if codeKecamatan == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	kecamatan, err := h.svc.GetVoterKecamatan(context.Background(), codeKecamatan)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	getVoterKecamatanResponse := kecamatan.ToGetVoterKecamatanResponse()

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get voter kelurahan success"),
		infrafiber.WithPayload(getVoterKecamatanResponse),
	).Send(ctx)

}

func (h handler) ListKecamatan(ctx *fiber.Ctx) error {

	kecamatans, err := h.svc.GetListKecamatan(context.Background())
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	kecamatanListResponse := NewKecamatanListResponseFromEntity(kecamatans)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list kecamatan success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(kecamatanListResponse),
	).Send(ctx)
}

func (h handler) GetListKelurahanFromKecamatan(ctx *fiber.Ctx) error {
	codeKecamatan := ctx.Params("code", "")
	if codeKecamatan == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	kelurahans, err := h.svc.GetListKelurahanFromKecamatan(context.Background(), codeKecamatan)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	kecamatanListResponse := NewKelurahanListResponseFromEntity(kelurahans)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list kelurahan success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(kecamatanListResponse),
	).Send(ctx)
}
