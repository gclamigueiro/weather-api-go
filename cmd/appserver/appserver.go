package appserver

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

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

	q := fmt.Sprintf(`%s,%s`, city, country)

	entry, inCache := cache.Get([]byte(q))

	var response []byte
	if inCache == nil {
		response = entry
		log.Println("Using cache for " + q)
	} else {
		result, err := weather.GetWeatherData(q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error calling weather endpoint"))
			return
		}
		response = []byte(result)
		cache.Set([]byte(q), response, expire)
		log.Println("Storing in cache " + q)
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
