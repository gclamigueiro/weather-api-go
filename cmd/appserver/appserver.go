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

func init() {

}

type Server struct {
	wheatherService weather.Service
	cache           *freecache.Cache
	expire          int
}

func NewServer() *Server {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	weatherService := weather.NewWeatherService()

	return &Server{
		cache:           cache,
		expire:          120,
		wheatherService: weatherService,
	}
}

func (s *Server) Start() {
	s.routes()
	log.Println("Listing for requests at http://localhost:8000/weather")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func (s *Server) routes() {
	http.HandleFunc("/weather", s.WeatherHandler)
}

func (s *Server) WeatherHandler(w http.ResponseWriter, req *http.Request) {

	queryParams := req.URL.Query()

	city, err := s.getQueryParamValue("city", queryParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("the city param is required"))
		return
	}

	country, err := s.getQueryParamValue("country", queryParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("the country param is required"))
		return
	}
	forecastday, _ := s.getQueryParamValue("forecastday", queryParams)
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

	entry, inCache := s.getFromCache(cacheKey)

	var response []byte
	if inCache == nil {
		response = entry
		log.Println("Using cache for " + cacheKey)
	} else {
		result, err := s.wheatherService.GetWeatherData(q, day)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error calling weather endpoint\n"))
			w.Write([]byte(err.Error()))
			log.Fatalln(err.Error())
			return
		}
		response = []byte(result)
		s.addToCache(cacheKey, response)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Check if a query param exist in request and return the value
func (s *Server) getQueryParamValue(param string, params url.Values) (string, error) {
	value, ok := params[param]
	if !ok || len(value) < 1 {
		return "", errors.New("parameter not found")
	}
	return value[0], nil
}

func (s *Server) getFromCache(key string) ([]byte, error) {
	value, err := s.cache.Get([]byte(key))
	return value, err
}

func (s *Server) addToCache(key string, response []byte) {
	log.Println("Storing in cache " + key)
	s.cache.Set([]byte(key), response, s.expire)
}
