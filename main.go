package main

import (
	"binance_testcase/router"
	"binance_testcase/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	apiDelivery "binance_testcase/api/delivery"
	apiRepository "binance_testcase/api/repository"
	apiUsecase "binance_testcase/api/usecase"
)

func main() {
	baseRouter := gin.New()
	baseRouter.Use(gin.Logger())
	baseRouter.Use(gin.Recovery())
	db, err := utils.InitPostgres()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	apiR := apiRepository.NewApiRepository(db)
	apiU := apiUsecase.NewApiUsecase(apiR)
	apiD := apiDelivery.NewApiDelivery(apiU)


	binanceRouter := baseRouter.Group("/")
	router.BinanceEndpoints(binanceRouter, apiD)
	
	server := &http.Server{
		Addr: ":8090",
		Handler: baseRouter,
		IdleTimeout: 10 * time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err = apiD.FillSymbols()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symbols, err := apiD.GetAllSymbols()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func(){
		for {
			apiD.GetPrices(symbols)
			time.Sleep(time.Minute * 1)
		}
	}()


	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("Coundn't start server")
			os.Exit(1)
		}
	}()
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	sig := <- sigChan
	fmt.Println("Graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	err = server.Shutdown(timeoutContext)
	if err != nil {
		fmt.Println("Graceful shutdown is not successful")
		os.Exit(1)
	}
}
