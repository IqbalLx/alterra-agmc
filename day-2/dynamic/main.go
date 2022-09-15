package main

import (
	"dynamic-crud/config"
	"dynamic-crud/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()

	e := echo.New()

	v1 := e.Group("/v1")
	routers.UserRouter(v1)

	e.Logger.Fatal(e.Start(":3000"))
}
