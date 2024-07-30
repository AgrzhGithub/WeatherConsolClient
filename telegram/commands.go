package telegram

import (
	"WeatherClient/client"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	MsgHello       = "Привет, "
	MsgHelp        = "Привет, Введите город в формате - Moscow: "
	MsgMissCommand = "Unknown command"
)

func helloMsg(b *bot.SendMessageParams, update *models.Update) string {
	b.Text = MsgHello + bot.EscapeMarkdown(update.Message.From.FirstName) + "!"
	return b.Text
}

func helpMsg(b *bot.SendMessageParams, update *models.Update) string {
	b.Text = MsgHelp
	return b.Text
}

func missMsg(b *bot.SendMessageParams, update *models.Update) string {
	b.Text = MsgMissCommand
	return b.Text
}

func weatherMsg(b *bot.SendMessageParams, update *models.Update) string {
	b.Text = "Temperature in " + update.Message.Text + ": " + client.GetWeatherData(update.Message.Text) + "°C"
	return b.Text
}
