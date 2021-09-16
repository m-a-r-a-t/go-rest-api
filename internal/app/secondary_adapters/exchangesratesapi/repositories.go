package exchangesratesapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type ExchangesRatesApiRepo struct {
	apiKey     string
	apiBaseUrl string
}

func NewExchangesRatesApiRepo() ExchangesRatesApiRepo {
	return ExchangesRatesApiRepo{
		apiKey:     "f04910d68152a0e43c04",
		apiBaseUrl: "https://free.currconv.com/api/v7/",
	}
}

func (p ExchangesRatesApiRepo) Convert(fromCurrency string,
	toCurrency string,
	amount float64,
) (float64, error) {
	var currentPrice float64
	var objmap map[string]interface{}
	log.Println("start")
	log.Println(fmt.Sprintf("%2.10f", amount))
	resp, err := http.Get(p.apiBaseUrl + "convert" +
		fmt.Sprintf(
			"?apiKey=%s&q=%s_%s&amount=%10.3f&compact=ultra",
			p.apiKey,
			fromCurrency,
			toCurrency,
			amount,
		),
	)
	if err != nil {
		log.Println(err)
		return currentPrice, err
	}
	defer resp.Body.Close()

	log.Println(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return currentPrice, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return currentPrice, err
	}

	err = json.Unmarshal(bytes, &objmap)
	if err != nil {
		return currentPrice, err
	}
	value := fmt.Sprintf("%f", objmap[fromCurrency+"_"+toCurrency])
	currentPrice, err = strconv.ParseFloat(value, 64)

	if err != nil {
		log.Println(err)
		return 0.0, err
	}

	return amount * currentPrice, nil
}
