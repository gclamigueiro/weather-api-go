package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gclamigueiro/weather-api-go/pkg/utility"
)

const weatherServer = "http://api.openweathermap.org/data/2.5/weather"
const apiKey = "1508a9a4840a5574c822d70ca2132032"

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

// Return the weather data of a specific city
func GetWeatherData(city, country string) (string, error) {
	q := fmt.Sprintf(`%s,%s`, city, country)
	endpoint := fmt.Sprintf(`%s?q=%s&appid=%s `, weatherServer, q, apiKey)

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
