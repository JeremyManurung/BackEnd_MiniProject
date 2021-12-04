package handler

import(
	"minipro/komentar"
	"minipro/helper"
	"github.com/labstack/echo"
	"minipro/user"
	"net/http"
)

type komentarHandler struct {
	service komentar.Service
}

func NewKomentarHandler(service komentar.Service) *komentarHandler {
	return &komentarHandler {service}
}

func (h *komentarHandler) CreateKomentar(echoContext echo.Context) error {
	var input komentar.CreateKomentarInput

	err := echoContext.Bind(&input)
	if err != nil {
		Response := helper.APIResponse("Gagal Buat Komentar", http.StatusUnprocessableEntity, "error", nil)
		return echoContext.JSON(http.StatusUnprocessableEntity, Response)
	}

	currentUser:= echoContext.Get("currentUser").(user.User)
	input.User = currentUser

	newKomentar, err := h.service.CreateKomentar(input)
	if err != nil {
		Response := helper.APIResponse("Gagal Membuat Komentar", http.StatusBadRequest, "error", nil)
		return echoContext.JSON(http.StatusOK, Response)
	}
	
	Response := helper.APIResponse("list bantuan", http.StatusOK, "succes", newKomentar)
	return echoContext.JSON(http.StatusOK, Response)
}

