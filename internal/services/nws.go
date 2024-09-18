package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/asilluron/weather-api/internal/models"
)

func GetWeatherForecast(latitude, longitude float64) (*models.Period, error) {

	header := "asilluron/weather-summary/1.0 (asilluron@no-reply.github.com)" // according to docs, this may change to an API key in the future
	pointURL := fmt.Sprintf("https://api.weather.gov/points/%.4f,%.4f", latitude, longitude)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", pointURL, nil)
	req.Header.Set("User-Agent", header)
	pointResp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting point data: %v", err)
	}
	defer pointResp.Body.Close()

	pointBody, err := io.ReadAll(pointResp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading point response: %v", err)
	}

	var point models.Point
	if err := json.Unmarshal(pointBody, &point); err != nil {
		return nil, fmt.Errorf("error unmarshaling point data: %v", err)
	}

	forecastReq, _ := http.NewRequest("GET", point.Properties.Forecast, nil)
	forecastReq.Header.Set("User-Agent", header)
	forecastResp, err := client.Do(forecastReq)
	if err != nil {
		return nil, fmt.Errorf("error getting forecast: %v", err)
	}
	defer forecastResp.Body.Close()

	forecastBody, err := io.ReadAll(forecastResp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading forecast response: %v", err)
	}

	var forecast models.Forecast
	if err := json.Unmarshal(forecastBody, &forecast); err != nil {
		return nil, fmt.Errorf("error unmarshaling forecast data: %v", err)
	}

	return &forecast.Properties.Periods[0], nil
}
