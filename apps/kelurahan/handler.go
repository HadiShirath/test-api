package kelurahan

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

func (h handler) GetKelurahanData(ctx *fiber.Ctx) error {
	codeKelurahan := ctx.Params("code", "")
	if codeKelurahan == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	kelurahan, err := h.svc.GetKeluharanData(context.Background(), codeKelurahan)
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

	getVoterTPSResponse := kelurahan.ToGetVoterKelurahanResponse()

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get voter kelurahan success"),
		infrafiber.WithPayload(getVoterTPSResponse),
	).Send(ctx)
}

func (h handler) GetListTPSFromKelurahan(ctx *fiber.Ctx) error {
	codeKelurahan := ctx.Params("code", "")
	if codeKelurahan == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	tpss, err := h.svc.GetListTPSFromKelurahan(context.Background(), codeKelurahan)
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

	kecamatanListResponse := NewTPSListResponseFromEntity(tpss)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list kelurahan success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(kecamatanListResponse),
	).Send(ctx)
}
