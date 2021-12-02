package main

import(
	"minipro/user"
	"minipro/helper"
	"minipro/handler"
	"minipro/auth"
	"minipro/bantuan"
	"minipro/transaksi"
	"minipro/pembayaran"
	"minipro/komentar"
	"strings"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	
)

var db *gorm.DB
func main(){
	dsn := "root:@tcp(host.docker.internal:3306)/backend?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if (err != nil) {
		panic(err)
	}

	userRepository := user.NewRepository(db)
	bantuanRepository := bantuan.NewRepository(db)
	transaksiRepository := transaksi.NewRepository(db)
	komentarRepository := komentar.NewRepository(db)

	bantuanService := bantuan.NewService(bantuanRepository)
	userService := user.NewService(userRepository)
	komentarService := komentar.NewService(komentarRepository)
	authService := auth.NewService()
	pembayaranService := pembayaran.NewService()
	transaksiService := transaksi.NewService(transaksiRepository, bantuanRepository, pembayaranService)

	userHandler := handler.NewUserHandler(userService, authService)
	bantuanHandler := handler.NewBantuanHandler(bantuanService)
	transaksiHandler := handler.NewTransaksiHandler(transaksiService)
	komentarHandler := handler.NewKomentarHandler(komentarService)
	r := echo.New()
	r.Static("/gambar", "./gambar")
	api :=r.Group("api/v1")

	api.POST("/check_email", userHandler.CheckEmail)
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/images",userHandler.UploadImg, authMiddleware(authService, userService))

	api.GET("/bantuans", bantuanHandler.GetBantuans)
	api.POST("/bantuans", bantuanHandler.CreateBantuan,authMiddleware(authService, userService))
	api.GET("/bantuans/:id",bantuanHandler.GetBantuan)

	api.GET("/bantuan/:id/transaksis", transaksiHandler.GetBantuanTransaksis,authMiddleware(authService, userService))
	api.GET("/transaksi", transaksiHandler.GetUserTransaksis,authMiddleware(authService, userService))
	api.POST("/transaksi", transaksiHandler.CreateTransaksi, authMiddleware(authService, userService))

	api.POST("/komentar", komentarHandler.CreateKomentar, authMiddleware(authService, userService))

	r.Start(":9000")
}

func authMiddleware(authService auth.Service, userService user.Service) echo.MiddlewareFunc{
return func (next echo.HandlerFunc) echo.HandlerFunc{
	return func(echoContext echo.Context) error{
		auth := ""
			for name, values := range echoContext.Request().Header {
				for _, value := range values {
					if name == "Authorization" {
						auth = value
					}
				}
			}

		if !strings.Contains(auth, "Bearer") {
			response := helper.APIResponse("Unauthorized contains", http.StatusUnauthorized, "error", nil)
			return echoContext.JSON(http.StatusUnauthorized, response)
		}

		tokenString := ""
		arrayToken := strings.Split(auth, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil{
			response := helper.APIResponse("Unauthorized validateToken", http.StatusUnauthorized, "error", nil)
			return echoContext.JSON(http.StatusUnauthorized, response)
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid{
			response := helper.APIResponse("Unauthorized claimtoken", http.StatusUnauthorized, "error", nil)
			return echoContext.JSON(http.StatusUnauthorized, response)
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil{
			response := helper.APIResponse("Unauthorized user", http.StatusUnauthorized, "error", nil)
			return echoContext.JSON(http.StatusUnauthorized, response)
		}

		echoContext.Set("currentUser", user)
		return next(echoContext)
}
}
}
