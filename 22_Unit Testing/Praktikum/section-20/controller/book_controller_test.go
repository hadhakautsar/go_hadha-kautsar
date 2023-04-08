package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"section-20/config"
	"section-20/model"
)

func TestGetBooksController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetBooksController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var res map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		if assert.NoError(t, err) {
			assert.Equal(t, "success get all books", res["message"])
			assert.NotNil(t, res["books"])
		}
	}
}

func TestGetBookController(t *testing.T) {
	e := echo.New()
	id := 1
	req := httptest.NewRequest(http.MethodGet, "/books/"+strconv.Itoa(id), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))

	book := model.Book{Title: "Test Book", Author: "Test Author", Publisher: "Test Publisher"}
	config.DB.Create(&book)

	if assert.NoError(t, GetBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var res map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		if assert.NoError(t, err) {
			assert.Equal(t, "success get book by id", res["message"])
			assert.NotNil(t, res["book"])
			assert.Equal(t, book.Title, res["book"].(map[string]interface{})["title"])
			assert.Equal(t, book.Author, res["book"].(map[string]interface{})["author"])
			assert.Equal(t, book.Publisher, res["book"].(map[string]interface{})["publisher"])
		}
	}

	config.DB.Delete(&book)
}

func TestCreateBookController(t *testing.T) {
	e := echo.New()
	book := model.Book{Title: "Test Book", Author: "Test Author", Publisher: "Test Publisher"}
	bookBytes, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(bookBytes))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var res map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		if assert.NoError(t, err) {
			assert.Equal(t, "success create new book", res["message"])
			assert.NotNil(t, res["book"])
			assert.Equal(t, book.Title, res["book"].(map[string]interface{})["title"])
			assert.Equal(t, book.Author, res["book"].(map[string]interface{})["author"])
			assert.Equal(t, book.Publisher, res["book"].(map[string]interface{})["publisher"])
		}
	}

	config.DB.Where("title = ?", book.Title).Delete(&model.Book{})
}

func TestUpdateBookController(t *testing.T) {
	e := echo.New()
	id := 1
	book := model.Book{Title: "Test Book", Author: "Test Author", Publisher: "Test Publisher"}
	config.DB.Create(&book)
	book.Title = "Updated Book"
	bookBytes, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPut, "/books/"+strconv.Itoa(id), bytes.NewBuffer(bookBytes))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	if assert.NoError(t, UpdateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var res map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		if assert.NoError(t, err) {
			assert.Equal(t, "success update book by id", res["message"])
			assert.NotNil(t, res["book"])
			assert.Equal(t, book.Title, res["book"].(map[string]interface{})["title"])
			assert.Equal(t, book.Author, res["book"].(map[string]interface{})["author"])
			assert.Equal(t, book.Publisher, res["book"].(map[string]interface{})["year"])
		}
	}

	config.DB.Delete(&book)
}
func TestDeleteBookController(t *testing.T) {
	e := echo.New()
	id := 1
	req := httptest.NewRequest(http.MethodDelete, "/books/"+strconv.Itoa(id), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))
	book := model.Book{Title: "Test Book", Author: "Test Author", Publisher: "Test Publisher"}
	config.DB.Create(&book)

	if assert.NoError(t, DeleteBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var res map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &res)
		if assert.NoError(t, err) {
			assert.Equal(t, "success delete book by id", res["message"])
		}
	}

	var count int64
	config.DB.Model(&model.Book{}).Where("id = ?", id).Count(&count)
	assert.Equal(t, int64(0), count)
}
