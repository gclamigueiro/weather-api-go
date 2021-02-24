package appserver

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/coocood/freecache"
	"github.com/gclamigueiro/weather-api-go/pkg/weather"
)

var cache *freecache.Cache

const expire = 120

func init() {
	cacheSize := 100 * 1024 * 1024
	cache = freecache.NewCache(cacheSize)
}

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

	forecastday, _ := getQueryParamValue("forecastday", queryParams)
	var day int

	if forecastday != "" {
		convertedDay, err := strconv.Atoi(forecastday)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid value for forecastday"))
			return
		}
		day = convertedDay

		// validate forecastday, must be a integer between 0 and 6
		if day < 0 || day > 6 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid value for forecastday"))
			return
		}
	}

	q := fmt.Sprintf(`%s,%s`, city, country)
	cacheKey := fmt.Sprintf(`%s,%d`, q, day)

	entry, inCache := cache.Get([]byte(cacheKey))

	var response []byte
	if inCache == nil {
		response = entry
		log.Println("Using cache for " + cacheKey)
	} else {
		result, err := weather.GetWeatherData(q, day)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error calling weather endpoint"))
			w.Write([]byte(err.Error()))
			log.Fatalln(err.Error())
			return
		}
		response = []byte(result)
		cache.Set([]byte(cacheKey), response, expire)
		log.Println("Storing in cache " + cacheKey)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func routes() {
	http.HandleFunc("/weather", WeatherHandler)
}

func Start() {
	routes()
	log.Println("Listing for requests at http://localhost:8000/weather")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
