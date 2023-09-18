package delivery

import (
	"binance_testcase/api"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"binance_testcase/models"

	"github.com/gin-gonic/gin"
)

type ApiDelivery struct {
	apiUsecase api.Usecase
}

func NewApiDelivery(apiUsecase api.Usecase) *ApiDelivery {
	return &ApiDelivery{
		apiUsecase: apiUsecase,
	}
}

func sendRequestToBinance(symbol string) (*models.MyStruct, error) {
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get("https://api.binance.com/api/v3/ticker?symbol="+symbol)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	var myStruct models.MyStruct
	err = json.NewDecoder(resp.Body).Decode(&myStruct)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &myStruct, nil
}

func sendRequestToBinancePair() (*models.Symbols, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("https://api.binance.com/api/v3/exchangeInfo")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	var myStruct models.Symbols
	err = json.NewDecoder(resp.Body).Decode(&myStruct)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &myStruct, nil
}

func (aD *ApiDelivery) FetchData(c *gin.Context) {
	symbolStruct, err := sendRequestToBinance("BNBBTC"); if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "err")
		return
	}
	
	
	c.JSON(200, symbolStruct)
}

func (aD *ApiDelivery) FillSymbols() (error) {
	symbolStruct, err := sendRequestToBinancePair(); if err != nil {
		fmt.Println(err)
		return err
	}

	err = aD.apiUsecase.FillSymbols(symbolStruct)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (aD *ApiDelivery) GetAllSymbols() ([]string, error) {
	symbols, err := aD.apiUsecase.GetAllSymbols()
	if err != nil {
		return []string{}, err
	}

	return symbols, nil
}


func (aD *ApiDelivery) GetPrices(symbols []string) (error) {
	goroutinesCount := 10
	wg := &sync.WaitGroup{}
	for i := 0; i < goroutinesCount; i++ {
		i := i
		wg.Add(1)
		go func(){
			for j := i; j < len(symbols); j+= goroutinesCount {
				result, err := sendRequestToBinance(symbols[j])
				if err != nil {
					continue
				}
				err = aD.apiUsecase.InsertPrice(result)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}
	wg.Wait()
	return nil
}

