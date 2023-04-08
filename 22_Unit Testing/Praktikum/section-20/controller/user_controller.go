package controller

import (
	"net/http"
	"strconv"

	"section-20/config"
	"section-20/middleware"
	"section-20/model"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []model.Users

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user model.Users

	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {

	user := model.Users{}
	c.Bind(&user)

	// validasi input user
	if user.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Name is required",
		})
	}
	if user.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Email is required",
		})
	}
	if user.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Password is required",
		})
	}

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT Token",
			"status":  err.Error(),
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&model.Users{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user model.Users

	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"user":    user,
	})

}

func LoginUsersController(c echo.Context) error {
	user := model.Users{}
	c.Bind(&user)

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail login user",
			"status":  err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"user":   user,
	})
}
