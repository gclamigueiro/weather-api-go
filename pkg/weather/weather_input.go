package weather

import "fmt"

// helper structures to get the data from the server
type weatherInput struct {
	Coord   coord
	Weather []weather
	Main    main
	Wind    wind
	Sys     sys
}

type coord struct {
	Lon float32
	Lat float32
}

func (t coord) String() string {
	// expected "[f, f]"
	return fmt.Sprintf("[%.2f, %.2f]", t.Lon, t.Lat)
}

type weather struct {
	Key         string `json:"main"`
	Description string
}

type main struct {
	Temp     float32
	Pressure int
	Humidity int
}

func (t main) PressureString() string {
	// expected 1027 hpa
	return fmt.Sprintf("%d hpa", t.Pressure)
}
func (t main) HumidityString() string {
	// expected Gentle breeze, 3.6 m/s, west-northwest
	return fmt.Sprintf("%d%%", t.Humidity)
}

type wind struct {
	Speed float32
	Deg   float32
}

func (t wind) String() string {
	// expected Gentle breeze, 3.6 m/s, west-northwest
	return fmt.Sprintf("%.1f m/s", t.Speed)
}

type sys struct {
	Sunrise int64
	Sunset  int64
}

// search for the key in Weather array
// it is used to retrieve Cloud Description
func (input weatherInput) GetWeatherDescription(key string) string {
	for _, w := range input.Weather {
		if w.Key == key {
			return w.Description
		}
	}
	return ""
}
