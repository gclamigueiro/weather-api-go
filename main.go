package main

import (
	"github.com/gclamigueiro/weather-api-go/cmd/appserver"
	_ "github.com/gclamigueiro/weather-api-go/pkg/config"
)

func main() {
	appserver.Start()
}
