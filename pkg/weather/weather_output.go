package weather

// structure that represents the response expected by the client
type weatherOutput struct {
	LocationName   string `json:"location_name"`
	Temperature    string `json:"temperature"`
	Wind           string `json:"wind"`
	Cloudiness     string `json:"cloudiness"`
	Pressure       string `json:"pressure"`
	Humidity       string `json:"humidity"`
	Sunrise        string `json:"sunrise"`
	Sunset         string `json:"sunset"`
	GeoCoordinates string `json:"geo_coordinates"`
	RequestedTime  string `json:"requested_time"`
	Forecast       string `json:"forecast"`
}
