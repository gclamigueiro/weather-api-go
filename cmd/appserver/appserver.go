package appserver

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/gclamigueiro/weather-api-go/pkg/weather"
)

// Check if a query param exist in request and return the value
func getQueryParamValue(param string, params url.Values) (string, error) {
	value, ok := params[param]
	if !ok || len(value) < 1 {
		return "", errors.New("parameter not found")
	}
	return value[0], nil
}

func WeatherHandler(w http.ResponseWriter, req *http.Request) {

	queryParams := req.URL.Query()

	city, err := getQueryParamValue("city", queryParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("the city param is required"))
		return
	}

	country, err := getQueryParamValue("country", queryParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("the country param is required"))
		return
	}

	response, err := weather.GetWeatherData(city, country)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error calling weather endpoint"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func routes() {
	http.HandleFunc("/weather", WeatherHandler)
}

func Start() {
	routes()
	log.Println("Listing for requests at http://localhost:8000/weather")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
