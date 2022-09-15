package controllers

import (
	"dynamic-crud/libs"
	"dynamic-crud/models"
	"dynamic-crud/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateUserController(context echo.Context) error {
	user := models.User{}
	err := json.NewDecoder(context.Request().Body).Decode(&user)
	if err != nil {
		return echo.ErrBadRequest
	}

	createdUser := libs.CreateUser(&user)
	return utils.ResponseHandler(context, http.StatusCreated, createdUser, "User created successfully")
}

func GetUserController(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	user, err := libs.GetUser(id)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return utils.ResponseHandler(context, http.StatusOK, user, "User fetched successfully")
}

func GetUsersController(context echo.Context) error {
	users, err := libs.GetUsers()
	if err != nil {
		return echo.ErrInternalServerError
	}
	return utils.ResponseHandler(context, http.StatusOK, users, "Users fetched successfully")
}

func UpdateUserController(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	user := models.User{}
	err = json.NewDecoder(context.Request().Body).Decode(&user)

	if err != nil {
		return echo.ErrBadRequest
	}

	isUserExists, err := libs.CheckExists(id)
	if !isUserExists || err != nil {
		return echo.ErrBadRequest
	}

	updatedUser, err := libs.UpdateUser(id, &user)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return utils.ResponseHandler(context, http.StatusOK, updatedUser, "User updated successfully")
}

func DeleteUserController(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	isUserExists, err := libs.CheckExists(id)
	if !isUserExists || err != nil {
		return echo.ErrBadRequest
	}

	libs.DeleteUser(id)
	return utils.ResponseHandler(context, http.StatusOK, nil, "User deleted successfully")
}
