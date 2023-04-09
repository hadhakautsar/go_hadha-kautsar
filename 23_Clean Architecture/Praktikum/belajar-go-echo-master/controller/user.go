package controller

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (c *UserController) GetAllUsers(ctx echo.Context) error {
	users, err := c.userUsecase.GetAllUsers()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"data": users,
	})
}

func (c *UserController) CreateUser(ctx echo.Context) error {
	var req dto.UserRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	user, err := c.userUsecase.CreateUser(req)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"data": user,
	})
}
