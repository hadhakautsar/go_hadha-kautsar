package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"section-20/config"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func TestGetUsersController(t *testing.T) {
	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		expected := `{"message":"success get all users","users":[]}` + "\n"
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestGetUserController(t *testing.T) {
	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		expected := `{"message":"success get user by id","user":[]}` + "\n"
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestCreateUserController(t *testing.T) {
	e := InitEcho()
	payload := map[string]string{
		"name":     "test",
		"email":    "test@gmail.com",
		"password": "123",
	}
	jsonPayload, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		expected := `{"message":"success create user","user":{"id":1,"name":"test","email":"test@gmail.com","password":"123","token":""}}` + "\n"
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestDeleteUserController(t *testing.T) {
	e := InitEcho()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		expected := `{"message":"success delete user by id"}` + "\n"
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestUpdateUserController(t *testing.T) {
	e := InitEcho()
	payload := map[string]string{
		"name": "Test User Updated",
	}
	jsonPayload, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(jsonPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		expected := `{"message":"success","user":{"id":1,"name":"Test User Updated","email":"testuser@example.com","password": "123","token":""}}`
		assert.JSONEq(t, expected, rec.Body.String())
	}
}

func TestLoginUsersController(t *testing.T) {
	e := InitEcho()
	payload := map[string]string{
		"email":    "testuser@example.com",
		"password": "pa55w0rd",
	}
	jsonPayload, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, LoginUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		expected := `{"status":"success login","user":{"id":1,"name":"Test User","email":"testuser@example.com","password":"123","token":""}}`
		assert.JSONEq(t, expected, rec.Body.String())
	}
}
