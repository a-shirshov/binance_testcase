package repository

import (
	"binance_testcase/models"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	createSymbol = `insert into symbol (symbol) values ($1);`
	getSymbols = `select symbol from symbol;`
	getSymbolID = `select id from symbol where symbol=$1;`
	insertPrice = `insert into bprice (symbol_id, price, timestamp) values ($1, $2, $3);`
)



type ApiRepository struct {
	db *sqlx.DB
}

func NewApiRepository(db *sqlx.DB) *ApiRepository {
	return &ApiRepository{
		db: db,
	}
}

func (aR *ApiRepository) FillSymbols(symbols *models.Symbols) (error) {
	for _, symbol := range symbols.Symbols {
		_, err := aR.db.Exec(createSymbol, symbol.Symbol)
		if err != nil {
			if strings.Contains(err.Error(), "(SQLSTATE 23505)") {
				continue
			}
			fmt.Println("Repo err", err)
			return err
		}
	}
	return nil
}

func (aR *ApiRepository) GetAllSymbols() ([]string, error) {
	var symbols []string
	err := aR.db.Select(&symbols, getSymbols)
	if err != nil {
		return []string{}, err
	}
	return symbols, nil
}

func (aR *ApiRepository) InsertPrice(symbol *models.MyStruct) (error) {
	var symbolID int
	err := aR.db.Get(&symbolID, getSymbolID, symbol.Ticker)
	if err != nil {
		return err
	}
	
	now := time.Now().Unix()
	_, err = aR.db.Exec(insertPrice, symbolID, symbol.Price, now)
	if err != nil {
		return err
	}
	return nil
}