package routers

import (
	"static-crud/controllers"

	"github.com/labstack/echo/v4"
)

func BookRouter(e *echo.Group) *echo.Group {
	bookRouter := e.Group("/book")

	bookRouter.POST("", controllers.CreateBookController)
	bookRouter.GET("", controllers.GetBooksController)
	bookRouter.GET("/:id", controllers.GetBookController)
	bookRouter.PUT("/:id", controllers.UpdateBookController)
	bookRouter.DELETE("/:id", controllers.DeleteBookController)

	return bookRouter
}
