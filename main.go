package main

import(
	"minipro/user"
	"minipro/handler"
	"minipro/auth"
	"minipro/bantuan"
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	
)

var db *gorm.DB
func main(){
	dsn := "root:@tcp(127.0.0.1:3306)/backend?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if (err != nil) {
		panic(err)
	}

	userRepository := user.NewRepository(db)
	bantuanRepository := bantuan.NewRepository(db)

	
	
	bantuanService := bantuan.NewService(bantuanRepository)
	userService := user.NewService(userRepository)
	authService := auth.NewService()


	userHandler := handler.NewUserHandler(userService, authService)
	bantuanHandler := handler.NewBantuanHandler(bantuanService)
	r := echo.New()
	
	api :=r.Group("api/v1")

	api.POST("/check_email", userHandler.CheckEmail)
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/images",userHandler.UploadImg)
	api.GET("/bantuans", bantuanHandler.GetBantuans)
	r.Start(":9000")
}

// func authMiddlerware(authService auth.Service, userService user.Service) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(echoContext echo.Context) error {
// 			//do the things
// 			authHeader := "Authorization"

// 		if !strings.Contains(authHeader, "Bearer") {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			return echoContext.JSON(http.StatusUnauthorized, response)
			
// 		}

// 		tokenString := ""
// 		arrayToken := strings.Split(authHeader, " ")
// 		if len(arrayToken) == 2 {
// 			tokenString = arrayToken[1]
// 		}

// 		token, err := authService.ValidateToken(tokenString)
// 		if err != nil {
// 			Response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			return echoContext.JSON(http.StatusUnauthorized, Response)
			
			
// 		}
		

// 		return next(echoContext)
// 	}
// }}
