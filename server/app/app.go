package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

const (
	CurrencyQuery = "vs_currency"
	MaxProc       = 4
	FilesDir      = "files"
)

type resultCurrencyJson struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	Symbol         string  `json:"symbol"`
	CurrentPrice   float64 `json:"current_price"`
	PriceChange24h float64 `json:"price_change_24h"`
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type App struct {
	apiPath string
	client  HttpClient
}

func NewApp(apiPath string, client HttpClient) *App {
	if client == nil {
		client = &http.Client{}
	}
	return &App{
		apiPath: apiPath,
		client:  client,
	}
}

func (a App) GetCryptoCurrencies(currencyStr string) (interface{}, error) {
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

type resultFile struct {
	FileName string `json:"file_name"`
	Found    bool   `json:"found"`
	Err      error  `json:"err"`
}

type resultJson struct {
	Results []resultFile `json:"results"`
}

func newResultJson() *resultJson {
	return &resultJson{
		Results: []resultFile{},
	}
}

func (rj *resultJson) add(res resultFile) {
	rj.Results = append(rj.Results, res)
}

type semaphore struct {
	sema chan struct{}
}

func newSemaphore() *semaphore {
	return &semaphore{
		sema: make(chan struct{}, MaxProc),
	}
}

func (s *semaphore) Acquire() {
	s.sema <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.sema
}

func (a App) FindHash(hashStr string) (interface{}, error) {
	var wg sync.WaitGroup
	sema := newSemaphore()
	resChan := make(chan resultFile, 10)
	resJson := newResultJson()
	for i := 1; i <= 10; i++ {
		// acquire resources
		sema.Acquire()
		wg.Add(1)

		// process
		go findInFile(fmt.Sprintf("file_%d.txt", i), hashStr, resChan, sema, &wg)
	}
	wg.Wait()
	close(resChan)
	for res := range resChan {
		resJson.add(res)
	}

	return resJson, nil
}

func findInFile(filename, hash string, resultChan chan resultFile, sema *semaphore, wg *sync.WaitGroup) {
	// release semaphore
	defer sema.Release()

	// release waitgroup
	defer wg.Done()

	res := resultFile{
		FileName: filename,
	}

	dat, err := os.ReadFile(strings.Join([]string{FilesDir, filename}, "/"))
	if err != nil {
		res.Err = err
		resultChan <- res
		return
	}
	res.Found = strings.Contains(string(dat), hash)
	resultChan <- res
	return
}
