package main

import (
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/matsubara0507/sample-go-bzlmod/front"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/static/*", echo.WrapHandler(front.EmbedFileServerHandler("/static/")))
	e.GET("/*", indexHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func indexHandler(ctx echo.Context) error {
	return ctx.HTMLBlob(http.StatusOK, front.EmbedIndexHtml())
}
