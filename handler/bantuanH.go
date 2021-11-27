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

	Response := helper.APIResponse("list bantuan", http.StatusOK, "succes", bantuans)
	return echoContext.JSON(http.StatusOK, Response)
}
