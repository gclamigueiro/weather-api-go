# weather-api-go

Example API in goland to retrieve the weather of a specific city. The data is obtained from <http://api.openweathermap.org>

## Run example

## create .env

Create .env with this values
```
weatherEnpoint = http://api.openweathermap.org/data/2.5/weather
forecastEndpoint = https://api.openweathermap.org/data/2.5/onecall
apiKey = {YourApiKey}
``` 

### get dependencies

``` 
go mod tidy 
```

### run

```
go run ./main.go 
```

## Run Test

```
go test ./test -v
```

## Example urls for testing

<http://localhost:8000/weather?city=Cali&country=co>

<http://localhost:8000/weather?city=Bogota&country=co&forecastday=2>
