package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type CoinGeckoResponse struct {
	Bitcoin struct {
		Usd float64 `json:"usd"`
	} `json:"bitcoin"`
	Ethereum struct {
		Usd float64 `json:"usd"`
	} `json:"ethereum"`
}

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

			msg := tu.Message(
				chatId,
				"Hello!",
			)
			bot.SendMessage(msg)

		}

	}, th.CommandEqual("start"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatId := tu.ID(update.Message.Chat.ID)
		go course("bitcoin", bot, chatId)
	}, th.CommandEqual("bitcoin"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatId := tu.ID(update.Message.Chat.ID)
		go course("ethereum", bot, chatId)
	}, th.CommandEqual("ethereum"))

	bh.Start()
}

func course(crypto string, bot *telego.Bot, Chatid telego.ChatID) {
	// Встановлення тайм-ауту для HTTP-клієнта
	client := &http.Client{Timeout: 10 * time.Second}

	// Відправка запиту до CoinGecko API
	resp, err := client.Get("https://api.coingecko.com/api/v3/simple/price?ids=" + crypto + "&vs_currencies=usd")
	if err != nil {
		log.Fatalf("Error fetching data from CoinGecko: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error response from CoinGecko: %s", resp.Status)
	}

	// Розбір JSON-відповіді
	var result CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}

	switch crypto {
	case "bitcoin":
		Result := fmt.Sprintf("Курс біткоіну на данний момент...", result.Bitcoin.Usd)
		bot.SendMessage(tu.Message(Chatid, Result))
	case "ethereum":
		Result := fmt.Sprintf("Курс ефіру на данний момент...", result.Ethereum.Usd)
		bot.SendMessage(tu.Message(Chatid, Result))
	default:
		bot.SendMessage(tu.Message(Chatid, "Ми не знайшли дані про цю крипто валюту! Перевірте правильність написання назви крипто валюти!"))
	}
}
