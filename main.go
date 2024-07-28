package main

import (
	Client "WeatherClient/client"
	"fmt"
)

func main() {

	fmt.Printf("Введите город в формате - 'Moscow': ")
	var ScanText string
	fmt.Scanf("%s", &ScanText)
	Client.GetWeatherData(ScanText)
}
