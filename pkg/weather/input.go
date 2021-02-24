package weather

import "fmt"

type Input interface {
	GetTemp() float32
	GetWind() wind
	GetWeatherDescription(key string) string
	GetPressure() pressure
	GetHumidity() humidity
	GetSunrise() int64
	GetSunset() int64
	GetGeoCoordinates() coord
}

type weather struct {
	Key         string `json:"main"`
	Description string
}

type coord struct {
	Lon float32
	Lat float32
}

type pressure int32
type humidity int32
type wind float32

func (t pressure) String() string {
	// expected 1027 hpa
	return fmt.Sprintf("%d hpa", t)
}

func (t humidity) String() string {
	return fmt.Sprintf("%d%%", t)
}

func (t wind) String() string {
	// expected Gentle breeze, 3.6 m/s, west-northwest
	return fmt.Sprintf("%.1f m/s", t)
}

func (t coord) String() string {
	return fmt.Sprintf("[%.2f, %.2f]", t.Lon, t.Lat)
}
