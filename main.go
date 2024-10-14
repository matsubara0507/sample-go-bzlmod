package main

import (
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed index.html
var indexHtmlFile []byte

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/*", indexHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func indexHandler(ctx echo.Context) error {
	return ctx.HTMLBlob(http.StatusOK, indexHtmlFile)
}
