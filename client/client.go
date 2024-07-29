package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	WeatherApi = "361a5a032b8d41f996b191849242707"
)

func GetWeatherData(city string) string {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=361a5a032b8d41f996b191849242707&q=%s&aqi=no", city)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather data", err)

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)

	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("Error unmarshalling weather data", err)

	}

	fmt.Printf("The current temperature in %s is %.2fC\n", weatherData.Location.Name, weatherData.Current.Temp)

	Temp := strconv.FormatFloat(weatherData.Current.Temp, 'f', 2, 64)
	return Temp
}
