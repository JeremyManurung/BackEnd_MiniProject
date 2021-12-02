package handler

import (
	"minipro/transaksi"
	"minipro/helper"
	
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
