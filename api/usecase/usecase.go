package usecase

import (
	"binance_testcase/api"
	"binance_testcase/models"
)

type ApiUsecase struct {
	apiRepository api.Repository
}

func NewApiUsecase(apiRepository api.Repository) *ApiUsecase {
	return &ApiUsecase{
		apiRepository: apiRepository,
	}
}

func (aU *ApiUsecase) FillSymbols(symbols *models.Symbols) (error) {
	return aU.apiRepository.FillSymbols(symbols)
}

func (aU *ApiUsecase) GetAllSymbols() ([]string, error) {
	return aU.apiRepository.GetAllSymbols()
}

func (aU *ApiUsecase) InsertPrice(symbol *models.MyStruct) (error) {
	return aU.apiRepository.InsertPrice(symbol)
}