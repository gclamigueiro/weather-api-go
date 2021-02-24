package appserver

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	weatherData "github.com/gclamigueiro/weather-api-go/pkg/weather_data"
)

// Check if a query param exist in request and return the value
func getQueryParamValue(param string, params url.Values) (string, error) {
	value, ok := params[param]
	if !ok || len(value) < 1 {
		return "", errors.New("parameter not found")
	}
	return value[0], nil
}

func weatherHandler(w http.ResponseWriter, req *http.Request) {

	queryParams := req.URL.Query()

	city, err := getQueryParamValue("city", queryParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("the city param  is required"))
		return
	}

	country, err := getQueryParamValue("country", queryParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("the country param  is required"))
		return
	}

	bodyString, err := weatherData.GetWeatherData(city, country)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error calling wheather endpoint"))
		return
	}

	//indentedContent, _ := json.MarshalIndent(bodyString, "", "    ")
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(bodyString))
}

func Start() {
	http.HandleFunc("/weather", weatherHandler)
	log.Println("Listing for requests at http://localhost:8000/weather")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
