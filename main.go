package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stock/stock"
)

func main() {
	r := mux.NewRouter()
	r.Path("/stock/{symbol}").Methods("GET").HandlerFunc(stock.GetStockBySymbol).Name("GetStockBySymbol")
	r.Path("/stock/{symbol}").Queries("stock_exchange", "{stock_exchange}").Methods("GET").HandlerFunc(stock.GetStockBySymbol).Name("GetStockBySymbol")
	if err := http.ListenAndServe(":9001", r); err != nil {
		log.Fatal(err)
	}
}
