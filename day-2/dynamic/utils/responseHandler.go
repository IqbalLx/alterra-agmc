package utils

import (
	"github.com/labstack/echo/v4"
)

func ResponseHandler(
	c echo.Context,
	status int,
	content interface{},
	message string,
) error {
	return c.JSON(status, map[string]interface{}{
		"message": message,
		"content": content,
	})
}
