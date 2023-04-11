package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/constants"
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	app := echo.New()
	app.GET("/users", userController.GetAllUsers)
	app.POST("/users", userController.CreateUser)
	app.POST("/login", userController.CreateUser)
	eJwt := app.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	app.GET("/users", userController.GetAllUsers)
	app.POST("/users", userController.CreateUser)
	app.Start(":8080")
}
