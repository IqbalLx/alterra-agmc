package main

import (
	"static-crud/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	v1 := e.Group("/v1")
	routers.BookRouter(v1)

	e.Logger.Fatal(e.Start(":3000"))
}
