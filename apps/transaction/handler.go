package transaction

import (
	"fmt"
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

func (h handler) CreateTransaction(ctx *fiber.Ctx) error {
	var req = CreateTransactionRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(err),
		).Send(ctx)
	}

	userPublicId := ctx.Locals("PUBLIC_ID")
	req.UserPublicId = fmt.Sprintf("%+v", userPublicId)

	if err := h.svc.CreateTransaction(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}
		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage(err.Error()),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("create transaction success"),
	).Send(ctx)
}

func (h handler) GetTransactionsByUserPublicId(ctx *fiber.Ctx) error {
	userPublicId := fmt.Sprintf("%+v", ctx.Locals("PUBLIC_ID"))

	trxs, err := h.svc.TransactionHistories(ctx.UserContext(), userPublicId)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage(err.Error()),
		).Send(ctx)
	}

	var response = []TransactionHistoryResponse{}

	for _, trx := range trxs {
		response = append(response, trx.ToTransactionHistoryResponse())
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(response),
		infrafiber.WithMessage("get transaction histories success"),
	).Send(ctx)
}
