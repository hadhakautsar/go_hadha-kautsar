package controller

import (
	"belajar-go-echo/config"
	"belajar-go-echo/dto"
	"belajar-go-echo/middleware"
	"belajar-go-echo/usecase"
	"net/http"

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

	token, err := middleware.CreateToken(int(user.ID), user.Email)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	response := map[string]interface{}{
		"data":  user,
		"token": token,
	}

	return ctx.JSON(http.StatusCreated, response)
}

func LoginUserController(c echo.Context) error {
	user := dto.UserRequest{}
	c.Bind(&user)

	db, err := config.ConnectDB()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to connect to database",
			"error":   err.Error(),
		})
	}

	err = db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to create token",
			"error":   err.Error(),
		})
	}

	userResponse := dto.UserResponse{
		ID:       uint(user.ID),
		Email:    user.Email,
		Password: user.Password,
		Token:    token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    userResponse,
	})
}
