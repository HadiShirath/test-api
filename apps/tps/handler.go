package tps

import (
	"context"
	"fmt"
	infrafiber "nbid-online-shop/infra/fiber"
	"nbid-online-shop/infra/response"
	"nbid-online-shop/internal/config"
	"nbid-online-shop/utility"
	"net/http"
	"os"
	"path/filepath"
	"time"

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
func (h handler) CreatePhoto(ctx *fiber.Ctx) error {

	file, errFile := ctx.FormFile("photo")
	if errFile != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid read file"),
			infrafiber.WithError(errFile),
			infrafiber.WithHttpCode(http.StatusNotFound),
		).Send(ctx)
	}

	filenameFromUser := file.Filename
	ext_photo := filepath.Ext(filenameFromUser)

	// TODO : CHECK EXTENSION FILE
	// file ekstensi yang diizinkan hanya jpg,jpeg,png
	// if ext_photo != ".jpg" && ext_photo != ".jpeg" && ext_photo != ".png" {
	// 	return infrafiber.NewResponse(
	// 		infrafiber.WithError(fiber.ErrBadRequest),
	// 		infrafiber.WithMessage("Invalid file extension. Allowed extensions: .jpg, .jpeg, .png"),
	// 		infrafiber.WithHttpCode(http.StatusBadRequest),
	// 	).Send(ctx)
	// }

	// get cookies user
	tokenUserData := ctx.Cookies("user")

	// validasi data yang sudah di generate token
	dataUser, err := utility.ValidateTokenData(tokenUserData, config.Cfg.App.Encryption.JWTSecret)
	if err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid process validate user"),
			infrafiber.WithError(err),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	// jika photo sebelumnya ada, maka hapus foto sebelumnya
	if dataUser.Photo != "" {
		filePath := filepath.Join("./public/images/", dataUser.Photo)
		os.Remove(filePath)
	}

	// nama file photo untuk simpan database
	timestamp := time.Now().Unix()
	filenameToDB := fmt.Sprintf("%+s-%+s-%+s-%+v%+s", dataUser.Kecamatan, dataUser.Kelurahan, dataUser.TPS, timestamp, ext_photo)

	// simpan file di direktori
	errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/images/%s", filenameToDB))
	if errSaveFile != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid process save file"),
			infrafiber.WithError(errSaveFile),
		).Send(ctx)
	}

	// create photo
	if err := h.svc.CreatePhoto(context.Background(), dataUser.ID, filenameToDB); err != nil {
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
		infrafiber.WithMessage("upload photo success"),
		infrafiber.WithHttpCode(http.StatusCreated),
	).Send(ctx)
}

func (h handler) TPSAdressDetail(ctx *fiber.Ctx) error {

	user_id := fmt.Sprintf("%+v", ctx.Locals("PUBLIC_ID"))

	tokenData, err := h.svc.TPSAdressDetail(context.Background(), user_id)
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

	// Set cookie dengan atribut yang sesuai
	ctx.Cookie(&fiber.Cookie{
		Name:     "user",                           // Nama cookie
		Value:    tokenData,                        // Nilai cookie
		Expires:  time.Now().Add(10 * time.Minute), // Waktu kedaluwarsa cookie
		Secure:   false,                            // Gunakan true jika aplikasi menggunakan HTTPS
		SameSite: "Lax",                            // "Lax", "Strict", atau "None" sesuai kebutuhan
	})

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get tps token"),
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithPayload(tokenData),
	).Send(ctx)
}

func (h handler) GetVoterTPS(ctx *fiber.Ctx) error {
	codeTPS := ctx.Params("code", "")
	if codeTPS == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	tps, err := h.svc.GetVoterTPS(context.Background(), codeTPS)
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

	getVoterTPSResponse := tps.ToGetVoterTPSResponse()

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get voter tps success"),
		infrafiber.WithPayload(getVoterTPSResponse),
	).Send(ctx)
}

func (h handler) GetAllVoterTPS(ctx *fiber.Ctx) error {

	tps, err := h.svc.GetAllVoterTPS(context.Background())

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

	getAllVoterTPSResponse := tps.ToGetAllVoterTPSResponse()

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("get all voter tps success"),
		infrafiber.WithPayload(getAllVoterTPSResponse),
	).Send(ctx)
}

func (h handler) GetTPSSaksiPagination(ctx *fiber.Ctx) error {
	var req = ListTPSSaksiRequestPayload{}

	if err := ctx.QueryParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("list tps invalid"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	tpss, err := h.svc.ListTPSSaksi(context.Background(), req)
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

	ListTPSSaksiResponse := NewTPSListSaksiResponseFromEntity(tpss)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("get list tps success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(ListTPSSaksiResponse),
		infrafiber.WithQuery(req.GenerateDefaultValue()),
	).Send(ctx)
}

func (h handler) GetListTPS(ctx *fiber.Ctx) error {

	tpss, err := h.svc.GetListTPS(context.Background())
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

func (h handler) UpdateVoteTPS(ctx *fiber.Ctx) error {
	var req = EditVoteTPSRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("list tps invalid"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	tpsId := ctx.Params("id", "")
	if tpsId == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid product"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	err := h.svc.EditVoteTPS(context.Background(), req, tpsId)
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

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage("update data tps success"),
	).Send(ctx)
}

func (h handler) UpdateVoteTPSByUser(ctx *fiber.Ctx) error {
	var req = EditVoteTPSBySaksiRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("data payload invalid"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	userId := ctx.Params("id", "")
	if userId == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid user id"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	data, err := h.svc.UpdateVoteTPSByUserId(context.Background(), req, userId)
	// err := h.svc.UpdateVoteTPSByUserId(context.Background(), req, userId)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid data tps"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	dataPayload := data.ToTpsDetailFromUpdateDataResponse()

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithMessage(fmt.Sprintf("Data Kecamatan %+s, Kelurahan %+s, %+s Berhasil ditambahkan", dataPayload.KecamatanName, dataPayload.KelurahanName, dataPayload.TpsName)),
	).Send(ctx)
}
