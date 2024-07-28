package client

type WeatherData struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`

	Current struct {
		Temp float64 `json:"temp_c"`
	} `json:"current"`
}
