package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	type CoinAPIResponse struct {
		Rate float64 `json:"rate"`
	}

	botToken := "7196274410:AAH07wSgzVqMJZRMUDpM4Gv2PMcZGlm7yVA"

	var URL string
	// api key
	apiKey := "9697AF10-3A86-4D0A-9DD7-A996B1BF8412"
	URL = "https://rest.coinapi.io/v1/assets/BTC" + apiKey
	// get info from api
	response, err := http.Get(URL)
	//ERROR(w, r, err, "Error!\n nil information in API")
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	// обробляєм помилку
	//ERROR(w, r, err, "Помилка читання тіла відповіді:")

	var coinAPIResp CoinAPIResponse
	err = json.Unmarshal(body, &coinAPIResp)
	// обробляєм помилку
	//ERROR(w, r, err, "Error!\n API LOCK!")
	Result := fmt.Sprintf("%.2f\n", coinAPIResp.Rate)

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

		msg := tu.Message(
			chatId,
			Result,
		)
		bot.SendMessage(msg)

	}, th.CommandEqual("btc"))

	bh.Start()
}
