package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	WeatherApi = "361a5a032b8d41f996b191849242707"
)

func GetWeatherData(city string) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=361a5a032b8d41f996b191849242707&q=%s&aqi=no", city)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather data", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)
		return
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("Error unmarshalling weather data", err)
		return
	}
	fmt.Printf("The current temperature in %s is %.2fC\n", weatherData.Location.Name, weatherData.Current.Temp)

}
