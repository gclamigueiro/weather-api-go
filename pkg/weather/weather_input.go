package weather

// helper structures to get the data from the server
type weatherInput struct {
	Coord   coord
	Weather []weather
	Main    main
	Wind    windObj
	Sys     sys
}

type main struct {
	Temp     float32
	Pressure int
	Humidity int
}

type windObj struct {
	Speed float32
	Deg   float32
}

type sys struct {
	Sunrise int64
	Sunset  int64
}

func (t weatherInput) GetTemp() float32 {
	return t.Main.Temp
}

func (t weatherInput) GetWind() wind {
	return wind(t.Wind.Speed)
}

func (t weatherInput) GetPressure() pressure {
	return pressure(t.Main.Pressure)
}

func (t weatherInput) GetHumidity() humidity {
	return humidity(t.Main.Humidity)
}

func (t weatherInput) GetSunrise() int64 {
	return t.Sys.Sunrise
}

func (t weatherInput) GetSunset() int64 {
	return t.Sys.Sunset
}

func (t weatherInput) GetGeoCoordinates() coord {
	return t.Coord
}

// search for the key in Weather array
// it is used to retrieve Cloud Description in case exist
func (input weatherInput) GetWeatherDescription(key string) string {
	for _, w := range input.Weather {
		if w.Key == key {
			return w.Description
		}
	}
	return ""
}
