package handler

import (
	"net/http"
	"minipro/helper"
	"os"
	"io"
	"fmt"
	"minipro/user"
	"github.com/labstack/echo"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func(h *userHandler) RegisterUser(echoContext echo.Context) error {
	var input user.RegisterUserInput

	err := echoContext.Bind(&input)
	if err != nil {
		Response := helper.APIResponse("Register Gagal", http.StatusUnprocessableEntity, "error", err.Error())
		return echoContext.JSON(http.StatusUnprocessableEntity, Response)
	}

	NewUser, err := h.userService.RegisterUser(input)
	if err != nil {
		Response := helper.APIResponse("Register Gagal", http.StatusBadRequest, "error", nil)
		return echoContext.JSON(http.StatusBadRequest, Response)
	}

	formatter := user.FormatUser(NewUser, "testes")
	Response := helper.APIResponse("Berhasil Register", http.StatusOK, "status", formatter)
	return echoContext.JSON(http.StatusOK, Response)
}

func(h *userHandler) Login(echoContext echo.Context) error {
	var input user.LoginInput

	err := echoContext.Bind(&input)
	if err != nil {
		Response := helper.APIResponse("Login Gagal", http.StatusUnprocessableEntity, "error", err.Error())
		return echoContext.JSON(http.StatusUnprocessableEntity, Response)
	}

	LoginUser, err := h.userService.Login(input)
	if err != nil {
		Response := helper.APIResponse("Login Gagal", http.StatusBadRequest, "error", nil)
		return echoContext.JSON(http.StatusBadRequest, Response)
	}

	formatter := user.FormatUser(LoginUser, "testes")
	Response := helper.APIResponse("Berhasil Login", http.StatusOK, "status", formatter)
	return echoContext.JSON(http.StatusOK, Response)
}

func(h *userHandler) CheckEmail(echoContext echo.Context) error {
	var input user.CheckEmailInput

	err := echoContext.Bind(&input)
	if err != nil {
		Response := helper.APIResponse("Email Gagal", http.StatusUnprocessableEntity, "error", err.Error())
		return echoContext.JSON(http.StatusUnprocessableEntity, Response)
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		Response := helper.APIResponse("Email Gagal", http.StatusBadRequest, "error", nil)
		return echoContext.JSON(http.StatusBadRequest, Response)
	}

	data :=  map [string] interface {}{
		"is_available" : IsEmailAvailable,
	}

	var metaMessage string

	if IsEmailAvailable {
		metaMessage = "Email Tidak Ada"
	}else{
		metaMessage = "Email Sudah Terdaftar"
	}

	Response := helper.APIResponse(metaMessage, http.StatusOK, "status", data)
	return echoContext.JSON(http.StatusOK, Response)
}

func(h *userHandler) UploadImg(echoContext echo.Context) error {
	file, err := echoContext.FormFile("img")
	if err != nil {
		data :=  map [string] interface {}{
		"is_uploaded" : false,
	}
		Response := helper.APIResponse("Login Gagal", http.StatusBadRequest, "error", data)
		return echoContext.JSON(http.StatusBadRequest, Response)
	}

	src, err := file.Open()
	if err != nil {
			data :=  map [string] interface {}{
		"is_uploaded" : false,
	}
		Response := helper.APIResponse("Login Gagal", http.StatusBadRequest, "error", data)
		return echoContext.JSON(http.StatusBadRequest, Response)
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
			data :=  map [string] interface {}{
		"is_uploaded" : false,
	}
		Response := helper.APIResponse("Login Gagal", http.StatusBadRequest, "error", data)
		return echoContext.JSON(http.StatusBadRequest, Response)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
			data :=  map [string] interface {}{
		"is_uploaded" : false,
	}
		Response := helper.APIResponse("Login Gagal", http.StatusBadRequest, "error", data)
		return echoContext.JSON(http.StatusBadRequest, Response)
	}

	userID := 2
	tes := fmt.Sprintf("gambar/%d-%s",userID,file.Filename)
	_, err = h.userService.SaveImg(userID, tes)
	if err != nil {
		data :=  map [string] interface {}{
		"is_uploaded" : false,
	}
		Response := helper.APIResponse("Gagal Upload", http.StatusBadRequest, "error", data)
		return echoContext.JSON(http.StatusBadRequest, Response)
	}

	data :=  map [string] interface {}{
		"is_uploaded" : true,
	}
	Response := helper.APIResponse("Berhasil Upload Gambar", http.StatusOK, "status", data)
	return echoContext.JSON(http.StatusOK, Response)
}





