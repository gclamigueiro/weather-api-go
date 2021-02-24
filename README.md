# weather-api-go

Example API in goland to retrieve the weather of a specific city. The data is consumed from <http://api.openweathermap.org>

## Run example

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

## Example url for testing

<http://localhost:8000/weather?city=Cali&country=co>

## Example of json received from openweathermap (only used fields)

```json
{
    "coord": {
        "lon": -74.0817,
        "lat": 4.6097
    },
    "weather": [
        {
            "main": "Clouds",
            "description": "broken clouds"
        }
    ],
    "main": {
        "temp": 291.15,
        "pressure": 1027,
        "humidity": 59
    },
    "wind": {
        "speed": 4.63,
        "deg": 50
    },
    "sys": {
        "sunrise": 1614164959,
        "sunset": 1614208201
    },
    "cod": 200
}
```

## Example of a json response from the weather-api-go

```json
{
    "location_name": "Cali,co",
    "temperature": "26.00°C",
    "wind": "2.1 m/s",
    "cloudiness": "scattered clouds",
    "pressure": "1016.00 hpa",
    "humidity": "65.0%",
    "sunrise": "06:18",
    "sunset": "18:20",
    "geo_coordinates": "[-76.52, 3.44]",
    "requested_time": "2021-02-24 10:58:48",
    "forecast": ""
}
```
