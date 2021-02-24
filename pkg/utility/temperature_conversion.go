package utility

import "fmt"

type Celsius float32

func (t Celsius) String() string {
	return fmt.Sprintf("%.2f °C", t)
}

func KelvinToCelsius(value float32) Celsius {
	celcius := value - 273.15
	return Celsius(celcius)
}
