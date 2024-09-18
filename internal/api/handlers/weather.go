package handlers

import (
	"fmt"
	"net/http"

	"github.com/asilluron/weather-api/internal/models"
	"github.com/asilluron/weather-api/internal/services"
	"github.com/labstack/echo/v4"
)

func HandleWeatherSummary(c echo.Context) error {
	coords := new(models.Coordinates)
	if err := c.Bind(coords); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input. Please provide lat and long."})
	}

	forecast, err := services.GetWeatherForecast(coords.Lat, coords.Long)
	if err != nil {
		// TODO: make a nice wrapper model for all errors
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Error fetching weather data: %v", err)})
	}

	response := models.WeatherResponse{
		Coordinates:        *coords,
		ShortForecast:      forecast.ShortForecast,
		TemperatureSummary: services.GetTemperatureSummary(forecast),
	}

	return c.JSON(http.StatusOK, response)
}
