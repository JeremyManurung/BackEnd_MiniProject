package handler

import (
	"minipro/transaksi"
	"minipro/helper"
	"minipro/user"
	"github.com/labstack/echo"
	"strconv"
	"net/http"
)

type transaksiHandler struct {
	service transaksi.Service
}

func NewTransaksiHandler(service transaksi.Service) *transaksiHandler {
	return &transaksiHandler{service}
}

func (h *transaksiHandler) GetBantuanTransaksis(echoContext echo.Context) error{
	id,_ := strconv.Atoi(echoContext.Param("id"))

	transaksis, err := h.service.GetTransaksisByBantuanID(id)
	if err != nil {
		Response := helper.APIResponse("Gagal menampilkan transaksi bantuan", http.StatusBadRequest, "error", nil)
		return echoContext.JSON(http.StatusBadRequest, Response)
	}

	Response := helper.APIResponse("berhasil menampilkan transaksi bantuan", http.StatusOK, "succes", transaksi.FormatBantuanTransaksis(transaksis))
	return echoContext.JSON(http.StatusOK, Response)
}

func (h *transaksiHandler) GetUserTransaksis(echoContext echo.Context) error{
	currentUser:= echoContext.Get("currentUser").(user.User)
	userID := currentUser.ID

	transaksis,err := h.service.GetTransaksisByUserID(userID)
	if err != nil {
		Response := helper.APIResponse("Gagal menampilkan transaksi user", http.StatusBadRequest, "error", nil)
		return echoContext.JSON(http.StatusBadRequest, Response)
	}

	Response := helper.APIResponse("berhasil menampilkan transaksi user", http.StatusOK, "succes", transaksi.FormatUserTransaksis(transaksis))
	return echoContext.JSON(http.StatusOK, Response)
}