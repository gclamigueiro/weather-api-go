package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gclamigueiro/weather-api-go/pkg/config"
	"github.com/gclamigueiro/weather-api-go/pkg/utility"
)

// transform the data received from openweathermap to the expected structure
func transformWeatherData(location string, input weatherInput) weatherOutput {

	output := weatherOutput{}
	output.LocationName = location
	output.Temperature = utility.KelvinToCelsius(input.Main.Temp).String()
	output.Wind = input.Wind.String()
	output.Cloudiness = input.GetWeatherDescription("Clouds")
	output.Pressure = input.Main.PressureString()
	output.Humidity = input.Main.HumidityString()

	sunriseTime := utility.TimeStampToDate(input.Sys.Sunrise)
	output.Sunrise = utility.GetFormattedTime(sunriseTime)
	sunsetTime := utility.TimeStampToDate(input.Sys.Sunset)
	output.Sunset = utility.GetFormattedTime(sunsetTime)

	output.GeoCoordinates = input.Coord.String()
	output.RequestedTime = time.Now().Format("2006-01-02 15:04:05")

	return output
}

func buildEndpoint(q string) string {
	weatherServer := config.GetConfig().WeatherServer
	apiKey := config.GetConfig().ApiKey
	endpoint := fmt.Sprintf(`%s?q=%s&appid=%s `, weatherServer, q, apiKey)
	return endpoint
}

// Return the weather data of a specific city
func GetWeatherData(q string) (string, error) {

	endpoint := buildEndpoint(q)

	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	weatherInput := weatherInput{}
	json.Unmarshal(bodyBytes, &weatherInput)

	// transform output in readable json
	weatherOutput := transformWeatherData(q, weatherInput)
	response, err := json.MarshalIndent(weatherOutput, "", " ")

	return string(response), err
}
