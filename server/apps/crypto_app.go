package apps

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	CurrencyQuery = "vs_currency"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type CryptoApp struct {
	apiPath string
	client  HttpClient
}

func NewCryptoApp(apiPath string, client HttpClient) *CryptoApp {
	if client == nil {
		client = &http.Client{}
	}
	return &CryptoApp{
		apiPath: apiPath,
		client:  client,
	}
}

type resultCurrencyJson struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	Symbol         string  `json:"symbol"`
	CurrentPrice   float64 `json:"current_price"`
	PriceChange24h float64 `json:"price_change_24h"`
}

func (a CryptoApp) GetCryptoCurrencies(currencyStr string) (interface{}, error) {
	req, err := http.NewRequest("GET",
		strings.Join([]string{a.apiPath, fmt.Sprintf("%s=%s", CurrencyQuery, currencyStr)}, "?"),
		nil)
	if err != nil {
		return nil, err
	}
	res, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	byteBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resJson []resultCurrencyJson
	err = json.Unmarshal(byteBody, &resJson)
	if err != nil {
		return nil, nil
	}

	return resJson, nil
}
