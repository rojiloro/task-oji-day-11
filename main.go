package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// static file from 'public' directory
	e.Static("/public", "public")

	e.GET("/hello", hai)

}

func hai(c echo.Context) error {
	return c.String(http.StatusOK, "haii dunia golang!!!, aku akan menaklukanmuuu....")
}