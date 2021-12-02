package handler

import (
	"minipro/bantuan"
	"minipro/helper"
	"minipro/user"
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

func(h *bantuanHandler) GetBantuan(echoContext echo.Context) error {

	id,_ := strconv.Atoi(echoContext.Param("id"))
	bantuanDetail, err := h.service.FindBantuanByID(id)
	if err != nil {
		Response := helper.APIResponse("Tidak ada detail bantuan", http.StatusBadRequest, "error", nil)
		return echoContext.JSON(http.StatusBadRequest, Response)
	}

	Response := helper.APIResponse("list bantuan", http.StatusOK, "succes", bantuan.FormatBantuanDetail(bantuanDetail))
	return echoContext.JSON(http.StatusOK, Response)
}



func (h *bantuanHandler) CreateBantuan(echoContext echo.Context) error {
	var input bantuan.CreateBantuanInput

	err := echoContext.Bind(&input)
	if err != nil {
		Response := helper.APIResponse("Gagal Buat Bantuan", http.StatusUnprocessableEntity, "error", nil)
		return echoContext.JSON(http.StatusUnprocessableEntity, Response)
	}

	currentUser:= echoContext.Get("currentUser").(user.User)
	input.User = currentUser

	newBantuan, err := h.service.CreateBantuan(input)
	if err != nil {
		Response := helper.APIResponse("Berhasil Membuat Bantuan", http.StatusOK, "success", bantuan.FormatBantuan(newBantuan))
		return echoContext.JSON(http.StatusOK, Response)
	}
	
	Response := helper.APIResponse("list bantuan", http.StatusOK, "succes", bantuan.FormatBantuan(newBantuan))
	return echoContext.JSON(http.StatusOK, Response)
}

// func (h *bantuanHandler) UpdateBantuan(echoContext echo.Context) error {
	
// 	inputID,_ := strconv.Atoi(echoContext.Param("id"))
// 	bantuan, err := h.service.FindBantuanByID(inputID)
// 	if err != nil {
// 		Response := helper.APIResponse("Tidak bisa update", http.StatusBadRequest, "error", nil)
// 		return echoContext.JSON(http.StatusBadRequest, Response)
// 	}


// 	var inputData bantuan.CreateBantuanInput
// 	err = echoContext.Bind(&inputData)
// 	if err != nil {
// 		Response := helper.APIResponse("Gagal updatenya gan", http.StatusUnprocessableEntity, "error", nil)
// 		return echoContext.JSON(http.StatusUnprocessableEntity, Response)
// 	}

// 	updatedBantuan, err := h.service.UpdateBantuan(inputID, inputData)
// 	if err != nil {
// 		Response := helper.APIResponse("Gagal updatenya gan", http.StatusUnprocessableEntity, "error", nil)
// 		return echoContext.JSON(http.StatusBadRequest, Response)
// 	}

// 	Response := helper.APIResponse("Bisa melakukan Update", http.StatusOK, "succes", bantuan.FormatBantuan(updatedBantuan))
// 	return echoContext.JSON(http.StatusOK, Response)
// }