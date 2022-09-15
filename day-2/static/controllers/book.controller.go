package controllers

import (
	"encoding/json"
	"net/http"
	"static-crud/libs"
	"static-crud/models"
	"static-crud/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateBookController(context echo.Context) error {
	book := models.Book{}
	err := json.NewDecoder(context.Request().Body).Decode(&book)
	if err != nil {
		return echo.ErrBadRequest
	}

	createdBook := libs.CreateBook(book)
	return utils.ResponseHandler(context, http.StatusCreated, createdBook, "book created successfully")
}

func GetBookController(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	book := libs.GetBook(id)
	return utils.ResponseHandler(context, http.StatusOK, book, "book fetched successfully")
}

func GetBooksController(context echo.Context) error {
	books := libs.GetBooks()
	return utils.ResponseHandler(context, http.StatusOK, books, "books fetched successfully")
}

func UpdateBookController(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	book := models.Book{}
	err = json.NewDecoder(context.Request().Body).Decode(&book)

	if err != nil {
		return echo.ErrBadRequest
	}

	updatedBook := libs.UpdateBook(id, book)
	return utils.ResponseHandler(context, http.StatusOK, updatedBook, "book updated successfully")
}

func DeleteBookController(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	book := libs.DeleteBook(id)
	return utils.ResponseHandler(context, http.StatusOK, book, "book deleted successfully")
}
