package main

import(
	"minipro/user"
	"minipro/handler"
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
	userService := user.NewService(userRepository)

	userService.SaveImg(1, "images/cekimg.jpg")

	userHandler := handler.NewUserHandler(userService)

	r := echo.New()
	
	api :=r.Group("api/v1")

	api.POST("/check_email", userHandler.CheckEmail)
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/images", userHandler.UploadImg)
	r.Start(":9000")
}

