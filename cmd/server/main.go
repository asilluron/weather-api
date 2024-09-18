package main

import (
	"github.com/asilluron/weather-api/internal/api/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/weather-summary", handlers.HandleWeatherSummary)
	// TODO: handle connection draining etc
	e.Logger.Fatal(e.Start(":8080"))
}
