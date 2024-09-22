package message

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

func (h handler) GetInboxList(ctx *fiber.Ctx) error {

	inboxs, err := h.svc.GetInboxList(context.Background())

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

	inboxListResponse := NewListInboxResponseFromEntity(inboxs)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list inbox success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(inboxListResponse),
	).Send(ctx)
}

func (h handler) GetOutboxList(ctx *fiber.Ctx) error {
	req := StatusMessageRequestPayload{}

	_ = ctx.QueryParser(&req)

	outboxs, err := h.svc.GetOutboxList(context.Background(), req)

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

	outboxListResponse := NewListOutboxResponseFromEntity(outboxs)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list inbox success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(outboxListResponse),
	).Send(ctx)
}

func (h handler) CreateMessage(ctx *fiber.Ctx) error {
	var req = CreateMessageRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	if err := h.svc.CreateMessage(context.Background(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("create message success"),
		infrafiber.WithHttpCode(http.StatusCreated),
	).Send(ctx)
}

func (h handler) UploadInbox(ctx *fiber.Ctx) error {
	var req = UploadInboxRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	if err := h.svc.UploadInbox(context.Background(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("upload inbox success"),
		infrafiber.WithHttpCode(http.StatusCreated),
	).Send(ctx)
}
