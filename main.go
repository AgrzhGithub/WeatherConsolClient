package main

import (
	"WeatherClient/telegram"
)

const token = "6314535674:AAHzBFX60JLsGwO-dn70bLweWA74xsbSv88"

func main() {

	//token := "6314535674:AAHzBFX60JLsGwO-dn70bLweWA74xsbSv88"
	//fmt.Printf("Введите город в формате - 'Moscow': ")
	//var ScanText string
	//fmt.Scanf("%s", &ScanText)
	//Client.GetWeatherData(ScanText)

	telegram.NewBot(token)
}
