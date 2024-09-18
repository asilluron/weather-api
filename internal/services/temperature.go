package services

import (
	"github.com/asilluron/weather-api/internal/models"
)

func GetTemperatureSummary(forecast *models.Period) string {
	if forecast.Temperature > 100 {
		return "scorching"
	}
	if forecast.Temperature > 89 {
		return "very hot"
	}
	if forecast.Temperature > 80 {
		return "hot"
	}
	if forecast.Temperature > 70 {
		return "moderate"
	}
	if forecast.Temperature > 60 {
		return "cool"
	}
	if forecast.Temperature > 45 {
		return "cold"
	}
	if forecast.Temperature > 32 {
		return "very cold"
	}
	return "freezing"
}
