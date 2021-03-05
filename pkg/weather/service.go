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

type Service interface {
	GetWeatherData(q string, day int) (string, error)
}

type WheaterService struct {
}

func NewWeatherService() Service {
	return &WheaterService{}
}

// Return the weather data of a specific city and day
func (s *WheaterService) GetWeatherData(q string, day int) (string, error) {

	endpoint := s.buildWeatherEndpoint(q)

	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var input Input

	weatherInput := &weatherInput{}
	json.Unmarshal(bodyBytes, weatherInput)

	// if the day is not zero, it is necessary to invoke the forecast endpoint
	// to obtain the specific day
	if day != 0 {
		lon := weatherInput.Coord.Lon
		lat := weatherInput.Coord.Lat
		dailyForecast, err := s.getForecastData(lon, lat, day)
		dailyForecast.Coord = weatherInput.Coord
		if err != nil {
			return "", err
		}
		input = dailyForecast
	} else {
		input = weatherInput
	}

	// transform output in readable json
	weatherOutput := s.createOutput(q, input)
	response, err := json.MarshalIndent(weatherOutput, "", " ")

	return string(response), err
}

// transform the data received from openweathermap to the expected structure
func (s *WheaterService) createOutput(location string, input Input) output {

	output := output{}
	output.LocationName = location
	output.Temperature = utility.KelvinToCelsius(input.GetTemp()).String()
	output.Wind = input.GetWind().String()
	output.Cloudiness = input.GetWeatherDescription("Clouds")
	output.Pressure = input.GetPressure().String()
	output.Humidity = input.GetHumidity().String()

	sunriseTime := utility.TimeStampToDate(input.GetSunrise())
	output.Sunrise = utility.GetFormattedTime(sunriseTime)
	sunsetTime := utility.TimeStampToDate(input.GetSunset())
	output.Sunset = utility.GetFormattedTime(sunsetTime)
	output.GeoCoordinates = input.GetGeoCoordinates().String()
	output.RequestedTime = time.Now().Format("2006-01-02 15:04:05")

	return output
}

func (s *WheaterService) buildWeatherEndpoint(q string) string {
	weatherEndpoint := config.GetConfig().WeatherEndpoint
	apiKey := config.GetConfig().ApiKey
	endpoint := fmt.Sprintf(`%s?q=%s&appid=%s `, weatherEndpoint, q, apiKey)
	return endpoint
}

func (s *WheaterService) buildForecastEndpoint(lon float32, lat float32) string {
	forecastEndpoint := config.GetConfig().ForecastEndpoint
	apiKey := config.GetConfig().ApiKey

	//https: //api.openweathermap.org/data/2.5/onecall?lat=33.441792&lon=-94.037689&exclude=hourly,minutely&appid={apikey}
	endpoint := fmt.Sprintf(`%s?lon=%f&lat=%f&exclude=hourly,minutely&appid=%s`, forecastEndpoint, lon, lat, apiKey)
	return endpoint
}

// Return the forecast data for a specific location and specific date
func (s *WheaterService) getForecastData(lon float32, lat float32, day int) (*forecastInput, error) {

	endpoint := s.buildForecastEndpoint(lon, lat)

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// transform output in readable json
	forecastInputList := forecastInputList{}
	json.Unmarshal(bodyBytes, &forecastInputList)

	// get specific date
	dayForecast, err := forecastInputList.GetForecastByDay(day)
	if err != nil {
		return nil, err
	}

	return dayForecast, nil
}
