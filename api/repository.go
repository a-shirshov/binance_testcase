package api

import "binance_testcase/models"

type Repository interface {
	FillSymbols(*models.Symbols) (error)
	GetAllSymbols() ([]string, error) 
	InsertPrice(symbol *models.MyStruct) (error) 
}


