package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := "7196274410:AAH07wSgzVqMJZRMUDpM4Gv2PMcZGlm7yVA"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		if update.Message != nil {
			chatId := tu.ID(update.Message.Chat.ID)

			keyboard := tu.Keyboard(
				tu.KeyboardRow(
					tu.KeyboardButton("Ethereum"),
				),
				tu.KeyboardRow(
					tu.KeyboardButton("Bitcoin"),
					tu.KeyboardButton("Ethereum"),
				),
				tu.KeyboardRow(
					tu.KeyboardButton("Bitcoin"),
					tu.KeyboardButton("Ethereum"),
				),
				tu.KeyboardRow(
					tu.KeyboardButton("Bitcoin"),
					tu.KeyboardButton("Ethereum"),
				),
				tu.KeyboardRow(
					tu.KeyboardButton("Bitcoin"),
					tu.KeyboardButton("Ethereum"),
				),
				tu.KeyboardRow(
					tu.KeyboardButton("Bitcoin"),
					tu.KeyboardButton("Ethereum"),
				),
				tu.KeyboardRow(
					tu.KeyboardButton("Bitcoin"),
					tu.KeyboardButton("Ethereum"),
				),
			).WithResizeKeyboard().WithInputFieldPlaceholder("Select something")

			msg := tu.Message(
				chatId,
				"Hello World",
			).WithReplyMarkup(keyboard)

			bot.SendMessage(msg)

			newmessfromuser := update.Message.Text

			switch newmessfromuser {
			case "Bitcoin":
				msg1 := tu.Message(
					chatId,
					"Hello World",
				)

				bot.SendMessage(msg1)
			}
		}

	}, th.CommandEqual("start"))

	bh.Start()
}
