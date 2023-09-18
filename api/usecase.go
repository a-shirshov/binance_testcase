package api

import "binance_testcase/models"

type Usecase interface {
	FillSymbols(*models.Symbols) (error)
	GetAllSymbols() ([]string, error) 
	InsertPrice(symbol *models.MyStruct) (error) 
}
