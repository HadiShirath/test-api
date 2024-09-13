package product

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

func (h handler) CreateProduct(ctx *fiber.Ctx) error {
	var req = CreateProductRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	if err := h.svc.CreateProduct(context.Background(), req); err != nil {
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
		infrafiber.WithMessage("create product success"),
		infrafiber.WithHttpCode(http.StatusCreated),
	).Send(ctx)
}

func (h handler) GetAllProduct(ctx *fiber.Ctx) error {
	var req = ListProductsRequestPayload{}

	if err := ctx.QueryParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("list product invalid"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	products, err := h.svc.ListProducts(context.Background(), req)
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

	ProductListResponse := NewProductListResponseFromEntity(products)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list product success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(ProductListResponse),
		infrafiber.WithQuery(req.GenerateDefaultValue()),
	).Send(ctx)
}

func (h handler) GetProductDetail(ctx *fiber.Ctx) error {
	sku := ctx.Params("sku", "")
	if sku == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	products, err := h.svc.ProductDetail(context.Background(), sku)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage("invalid product"),
		).Send(ctx)
	}

	ProductDetailResponse := products.ToProductDetailResponse()

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get product success"),
		infrafiber.WithPayload(ProductDetailResponse),
	).Send(ctx)
}
