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

		/*u := make(Map[string]string)

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tShould decode the response.", ballotX)
		}
		t.Log("\tShould decode the response.", checkMark)

		if u.Name == "Bill" {
			t.Log("\tShould have a Name.", checkMark)
		} else {
			t.Error("\tShould have a Name.", ballotX, u.Name)
		}

		if u.Email == "bill@ardanstudios.com" {
			t.Log("\tShould have an Email.", checkMark)
		} else {
			t.Error("\tShould have an Email.", ballotX, u.Email)
		}*/
	}
}
