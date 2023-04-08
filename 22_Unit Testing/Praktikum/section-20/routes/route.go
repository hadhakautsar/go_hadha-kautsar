package routes

import (
	"section-20/constants"
	"section-20/controller"
	m "section-20/middleware"

	echomid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	m.LogMiddleware(e)
	// Route / to handler function
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)
	e.POST("/users/login", controller.LoginUsersController)
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)
	e.POST("/books", controller.CreateBookController)
	e.PUT("/books/:id", controller.UpdateBookController)
	e.DELETE("/books/:id", controller.DeleteBookController)

	eJwt := e.Group("/jwt")
	eJwt.Use(echomid.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users", controller.GetUserController)
	eJwt.GET("/users/:id", controller.GetUserController)
	eJwt.PUT("/users/:id", controller.UpdateUserController)
	eJwt.DELETE("/users/:id", controller.DeleteUserController)
	eJwt.GET("/books", controller.GetBooksController)
	eJwt.GET("/books/:id", controller.GetBookController)
	eJwt.POST("/books", controller.CreateBookController)
	eJwt.PUT("/books/:id", controller.UpdateBookController)
	eJwt.DELETE("/books/:id", controller.DeleteBookController)

	return e
}
