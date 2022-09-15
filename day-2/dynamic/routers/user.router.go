package routers

import (
	"dynamic-crud/controllers"

	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Group) *echo.Group {
	userRouter := e.Group("/user")

	userRouter.POST("", controllers.CreateUserController)
	userRouter.GET("", controllers.GetUsersController)
	userRouter.GET("/:id", controllers.GetUserController)
	userRouter.PUT("/:id", controllers.UpdateUserController)
	userRouter.DELETE("/:id", controllers.DeleteUserController)

	return userRouter
}
