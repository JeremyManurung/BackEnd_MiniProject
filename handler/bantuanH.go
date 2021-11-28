package handler

import (
	"minipro/bantuan"
	"minipro/helper"
	"strconv"
	"net/http"
	"github.com/labstack/echo"
)

type bantuanHandler struct {
	service bantuan.Service
}

func NewBantuanHandler(service bantuan.Service) *bantuanHandler {
	return &bantuanHandler {service}
}

func(h *bantuanHandler) GetBantuans(echoContext echo.Context) error{
	userID, _ := strconv.Atoi(echoContext.Param("user_id"))

	bantuans, err := h.service.FindBantuans(userID)
	if err != nil {
		Response := helper.APIResponse("Tidak dapat membuat bantuan", http.StatusBadRequest, "error", nil)
		return echoContext.JSON(http.StatusBadRequest, Response)
	}

	Response := helper.APIResponse("list bantuan", http.StatusOK, "succes", bantuan.FormatBantuans(bantuans))
	return echoContext.JSON(http.StatusOK, Response)
}

// func (h *bantuanHandler) CreateBantuan(echoContext echo.Context) error {
// 	var input bantuan.CreateBantuanInput

// 	err := echoContext.Bind(&input)
// 	if err != nil {
// 		Response := helper.APIResponse("Gagal Buat Bantuan", http.StatusUnprocessableEntity, "error", nil)
// 		return echoContext.JSON(http.StatusUnprocessableEntity, Response)
// 	}

// 	userID := 1
// 	newBantuan, err := h.service.CreateBantuan(userID)
// 	if err != nil {
// 		Response := helper.APIResponse("Berhasil Membuat Bantuan", http.StatusOK, "success", bantuan.FormatBantuan(newBantuan))
// 		return echoContext.JSON(http.StatusOK, Response)
// 	}
	
// 	Response := helper.APIResponse("list bantuan", http.StatusOK, "succes", bantuan.FormatBantuans(bantuans))
// 	return echoContext.JSON(http.StatusOK, Response)
// }
