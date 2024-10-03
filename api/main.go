package main

import (
	"net/http"
	"wheater/cmd/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{"Content-Type", "Accept"},
	}))

	e.GET("/getweather/:city", handlers.GetWheater) //TEMPLATE RENDER
	e.GET("/health", handlers.Health)

	e.Logger.Fatal(e.Start(":8002"))
}
