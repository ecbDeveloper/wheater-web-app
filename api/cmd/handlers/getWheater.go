package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"wheater/cmd/models"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func GetWheater(c echo.Context) error {
	var weatherData models.WeatherData

	if err := godotenv.Load(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	apiKey := os.Getenv("API_KEY")

	city := c.Param("city")

	spacedCity := url.QueryEscape(city)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&lang=pt_br&appid=%s&units=metric", spacedCity, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Failed to make the request", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read the response body", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	if err = json.Unmarshal(body, &weatherData); err != nil {
		log.Println("Failed to decode response body")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if weatherData.Name == "" {
		log.Println("City not found")
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, weatherData)
}
