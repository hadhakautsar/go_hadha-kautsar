package routes

import (
	"section-20/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)
	e.POST("/books", controller.CreateBookController)
	e.PUT("/books/:id", controller.UpdateBookController)
	e.DELETE("/books/:id", controller.DeleteBookController)
	return e
}
