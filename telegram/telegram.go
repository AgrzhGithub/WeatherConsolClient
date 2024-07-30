package telegram

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"os"
	"os/signal"
	"strings"
)

func NewBot(token string) {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/hello", bot.MatchTypeExact, handler)

	b.Start(ctx)

}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   checkInput(update),
	})
}

func checkInput(update *models.Update) string {
	b := &bot.SendMessageParams{}
	str := strings.Fields(update.Message.Text)

	if update.Message != nil && len(str) == 1 {
		switch update.Message.Text {
		case "/start":
			helloMsg(b, update)
		case "/help":
			helpMsg(b, update)
		default:
			if strings.HasPrefix(update.Message.Text, "/") {
				missMsg(b, update)
			} else {
				weatherMsg(b, update)
			}
		}
	} else {
		missMsg(b, update)
	}
	return b.Text
}
