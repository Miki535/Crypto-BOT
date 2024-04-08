package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type CoinAPIResponse struct {
	Rate float64 `json:"rate"`
}

func osn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var URL string
		var cryptovalue string
		// парсим інформацію з html сторінки і обробляєм її з верхнього регістру
		cprypto := r.FormValue("Cryptoval")
		converted := strings.Title(cprypto)
		// перевірка того що написав користувач
		switch converted {
		case "Bitcoin":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"
			cryptovalue = "bitcoin"
		case "Ethereum":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd"
			cryptovalue = "ethereum"
		case "Fantom":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=fantom&vs_currencies=usd"
			cryptovalue = "fantom"
		case "BNB":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=bnb&vs_currencies=usd"
			cryptovalue = "bnb"
		case "Solana":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=solana&vs_currencies=usd"
			cryptovalue = "solana"
		case "XRP":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=xrp&vs_currencies=usd"
			cryptovalue = "xrp"
		case "Dogecoin":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=dogecoin&vs_currencies=usd"
			cryptovalue = "dogecoin"
		case "Polkadot":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=polkadot&vs_currencies=usd"
			cryptovalue = "polkadot"
		case "Chainlink":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=chainlink&vs_currencies=usd"
			cryptovalue = "chainlink"
		case "Toncoin":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=toncoin&vs_currencies=usd"
			cryptovalue = "toncoin"
		case "Uniswap":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=uniswap&vs_currencies=usd"
			cryptovalue = "uniswap"
		case "Litecoin":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=litecoin&vs_currencies=usd"
			cryptovalue = "litecoin"
		case "Pepe":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=pepe&vs_currencies=usd"
			cryptovalue = "pepe"
		case "Gnosis":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=gnosis&vs_currencies=usd"
			cryptovalue = "gnosis"
		case "Neo":
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=neo&vs_currencies=usd"
			cryptovalue = "neo"
		default:

		}

		response, err := http.Get(URL)
		ERROR(w, r, err, "Error!\n nil information in API")
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, "Помилка читання тіла відповіді:", http.StatusInternalServerError)
			return
		}

		var data map[string]map[string]float64
		if err := json.Unmarshal(body, &data); err != nil {
			ApiUse(w, r)
			return
		}
		btcUsd := data[cryptovalue]["usd"]

		Data := struct {
			Result float64
		}{
			Result: btcUsd,
		}

		tpl.Execute(w, Data)
		return
	}
	tpl.Execute(w, nil)
}

func ApiUse(w http.ResponseWriter, r *http.Request) {
	cprypto := r.FormValue("Cryptoval")
	converted := strings.Title(cprypto)
	var URL string
	// api key
	apiKey := "9697AF10-3A86-4D0A-9DD7-A996B1BF8412"
	// перевірка того що написав користувач
	switch converted {
	case "Bitcoin":
		URL = "https://rest.coinapi.io/v1/assets/BTC" + apiKey
	case "Ethereum":
		URL = "https://rest.coinapi.io/v1/assets/ETH" + apiKey
	case "Fantom":
		URL = "https://rest.coinapi.io/v1/assets/FMT" + apiKey
	default:
		http.Error(w, "Невірний ввід для криптовалюти", http.StatusBadRequest)
		return
	}
	// дістаєм інформацію з api
	response, err := http.Get(URL)
	ERROR(w, r, err, "Error!\n nil information in API")
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	// обробляєм помилку
	ERROR(w, r, err, "Помилка читання тіла відповіді:")

	var coinAPIResp CoinAPIResponse
	err = json.Unmarshal(body, &coinAPIResp)
	// обробляєм помилку
	ERROR(w, r, err, "Error!\n API LOCK!")
	new := fmt.Sprintf("%.2f\n", coinAPIResp.Rate)
	// data з інформацією
	Data := struct {
		Result string
	}{
		Result: new,
	}
	tpl.Execute(w, Data)
	return
}

func team(w http.ResponseWriter, r *http.Request) {
	tplp.Execute(w, nil)
}

// func ERROR фунція що обробляє помилку і відправляє на html сторінку
func ERROR(w http.ResponseWriter, r *http.Request, err error, text string) {
	if err != nil {
		http.Error(w, text, http.StatusBadRequest)
		return
	}
}
