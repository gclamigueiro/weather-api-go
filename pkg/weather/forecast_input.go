package weather

import (
	"errors"
)

// helper structures to get forecast data from the server

type forecastInput struct {
	Temp      temp
	Pressure  int
	Humidity  int
	WindSpeed float32 `json:"wind_speed"`
	WindDeg   float32
	Sunrise   int64
	Sunset    int64
	Weather   []weather
	Coord     coord
}

type temp struct {
	Day float32
}

type forecastInputList struct {
	Daily []*forecastInput
}

func (t forecastInput) GetTemp() float32 {
	return t.Temp.Day
}

func (t forecastInput) GetWind() wind {
	// expected Gentle breeze, 3.6 m/s, west-northwest
	return wind(t.WindSpeed)
}

func (t forecastInput) GetPressure() pressure {
	return pressure(t.Pressure)
}

func (t forecastInput) GetHumidity() humidity {
	return humidity(t.Humidity)
}

func (t forecastInput) GetSunrise() int64 {
	return t.Sunrise
}

func (t forecastInput) GetSunset() int64 {
	return t.Sunset
}

func (t forecastInput) GetGeoCoordinates() coord {
	return t.Coord
}

// search for the key in Weather array
// it is used to retrieve Cloud Description in case exist
func (input forecastInput) GetWeatherDescription(key string) string {
	for _, w := range input.Weather {
		if w.Key == key {
			return w.Description
		}
	}
	return ""
}

// search specific day in forecastInputList
func (input forecastInputList) GetForecastByDay(day int) (*forecastInput, error) {
	if day < len(input.Daily) {
		return input.Daily[day], nil
	}
	return nil, errors.New("day not found")
}
