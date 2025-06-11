package main

/*
To Do:
- Cleanup + DRY for API calls here
- More pretty UI
*/

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"encoding/json"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

type WeatherResponse struct {
    Location Location `json:"location"`
    Current  Current  `json:"current"`
}

type Location struct {
    Name           string  `json:"name"`
    Region         string  `json:"region"`
    Country        string  `json:"country"`
    Lat            float64 `json:"lat"`
    Lon            float64 `json:"lon"`
    TzID           string  `json:"tz_id"`
    LocaltimeEpoch int64   `json:"localtime_epoch"`
    Localtime      string  `json:"localtime"`
}

type Current struct {
    LastUpdatedEpoch int64     `json:"last_updated_epoch"`
    LastUpdated      string    `json:"last_updated"`
    TempC            float64   `json:"temp_c"`
    TempF            float64   `json:"temp_f"`
    IsDay            int       `json:"is_day"`
    Condition        Condition `json:"condition"`
    WindMph          float64   `json:"wind_mph"`
    WindKph          float64   `json:"wind_kph"`
    WindDegree       int       `json:"wind_degree"`
    WindDir          string    `json:"wind_dir"`
    PressureMb       float64   `json:"pressure_mb"`
    PressureIn       float64   `json:"pressure_in"`
    PrecipMm         float64   `json:"precip_mm"`
    PrecipIn         float64   `json:"precip_in"`
    Humidity         int       `json:"humidity"`
    Cloud            int       `json:"cloud"`
    FeelslikeC       float64   `json:"feelslike_c"`
    FeelslikeF       float64   `json:"feelslike_f"`
    WindchillC       float64   `json:"windchill_c"`
    WindchillF       float64   `json:"windchill_f"`
    HeatindexC       float64   `json:"heatindex_c"`
    HeatindexF       float64   `json:"heatindex_f"`
    DewpointC        float64   `json:"dewpoint_c"`
    DewpointF        float64   `json:"dewpoint_f"`
    VisKm            float64   `json:"vis_km"`
    VisMiles         float64   `json:"vis_miles"`
    UV               float64   `json:"uv"`
    GustMph          float64   `json:"gust_mph"`
    GustKph          float64   `json:"gust_kph"`
}

type Condition struct {
    Text string `json:"text"`
    Icon string `json:"icon"`
    Code int    `json:"code"`
}

func fetchWeather(c *gin.Context) {
	apiKey := os.Getenv("API_KEY")

	// fetch client IP from context
	ip, exists := c.Get("clientIp")
	if ! exists {
		log.Fatal("no IP address set")
	}

	// overwrite IP address (my local here in Austin)
	if ip == "127.0.0.1" {
		ip = "108.65.112.223"
	}

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, ip)
	fmt.Println(url)

	var weather WeatherResponse

	resp, err := http.Get(url + apiKey)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	err = json.Unmarshal([]byte(string(body)), &weather)
    if err != nil {
        log.Fatal(err)
    }

	c.IndentedJSON(http.StatusOK, weather)
}

func searchWeather(c *gin.Context) {
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, c.Param("query"))

	fmt.Println(url)

	var weather WeatherResponse

	resp, err := http.Get(url + apiKey)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	err = json.Unmarshal([]byte(string(body)), &weather)
    if err != nil {
        log.Fatal(err)
    }

	c.IndentedJSON(http.StatusOK, weather)
}

func fetchClientIp (c *gin.Context) {
	c.Set("clientIp", c.ClientIP())
    c.Next()
}

func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	// setup gin router
	router := gin.Default()
	router.Use(fetchClientIp)
    router.GET("/", fetchWeather)
	router.GET("/search:query", searchWeather)
    router.Run(fmt.Sprintf("localhost:%s", port))
}