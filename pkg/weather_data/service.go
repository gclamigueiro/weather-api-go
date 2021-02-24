package weather_data

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const wheaterServer = "http://api.openweathermap.org/data/2.5/weather"
const apiKey = "1508a9a4840a5574c822d70ca2132032"

// Return the weather data of a specific city
func GetWeatherData(city, country string) (string, error) {
	q := fmt.Sprintf(`%s,%s `, city, country)
	endpoint := fmt.Sprintf(`%s?q=%s&appid=%s `, wheaterServer, q, apiKey)

	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(bodyBytes)

	return bodyString, err
}
