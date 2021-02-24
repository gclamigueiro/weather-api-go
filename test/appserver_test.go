package appserver_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gclamigueiro/weather-api-go/cmd/appserver"
	"github.com/joho/godotenv"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	godotenv.Load("../.env")
}

// TestSendJSON testing the sendjson internal endpoint.
func TestWeatherHandler(t *testing.T) {
	t.Log("Testing WeatherHandler Endpoint .")
	{

		// TODO create Table tests

		city := "Cali"
		country := "co"

		endpoint := fmt.Sprintf(`weather?city=%s&country=%s `, city, country)

		req, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			t.Fatal("\tShould be able to create a request.",
				ballotX, err)
		}
		t.Log("\tShould be able to create a request.",
			checkMark)

		rw := httptest.NewRecorder()
		handler := http.HandlerFunc(appserver.WeatherHandler)
		handler.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
		}
		t.Log("\tShould receive \"200\"", checkMark)

		// TODO validate response

	}
}
