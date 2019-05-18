package stock

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
	"github.com/stock/config"
	"github.com/stock/model"
)

func GetStockBySymbol(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]
	if len(symbol) < 1 {
		model.WriteErrorResponse(w, http.StatusInternalServerError, "Error! The requested stock(s) could not be found.")
		return
	}

	stockExchange := r.FormValue("stock_exchange")
	url := config.STOCK_URL
	url += "symbol=" + symbol
	url += "&api_token=" + config.API_USER_TOKEN

	response, err := http.Get(url)
	if err != nil {
		model.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		model.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var output model.APIResponse
	var re model.Result
	var res []model.Result
	err = json.Unmarshal(data, &output)
	if err != nil {
		model.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(output.Data) < 1 {
		model.WriteErrorResponse(w, http.StatusInternalServerError, "Error! The requested stock(s) could not be found.")
		return
	}

	for _, d := range output.Data {
		re.Symbol = d["symbol"]
		re.Name = d["name"]
		re.CloseYesterday = d["close_yesterday"]
		re.Currency = d["currency"]
		re.GmtOffset = d["gmt_offset"]
		re.LastTradeTime = d["last_trade_time"]
		re.MarketCap = d["market_cap"]
		re.Price = d["price"]
		re.Timezone = d["timezone"]
		re.TimezoneName = d["timezone_name"]
		re.Volume = d["volume"]
		re.StockExchange = d["stock_exchange_short"]
		res = append(res, re)
	}

	if len(stockExchange) < 1 {
		stockExchange = config.DEFAULT_STOCK_EXCHANGE
	}

	res = filterResultByStockExchange(res, stockExchange)
	model.WriteOKResponse(w, res)
}

func filterResultByStockExchange(data []model.Result, stockExchange string) []model.Result {
	arrStockExchange := strings.Split(stockExchange, ",")
	sort.Strings(arrStockExchange)
	returnStockExchange := []model.Result{}
	for _, stockExchange := range data {
		if (len(arrStockExchange)-1) != sort.SearchStrings(arrStockExchange, stockExchange.StockExchange) || arrStockExchange[len(arrStockExchange)-1] == stockExchange.StockExchange {
			returnStockExchange = append(returnStockExchange, stockExchange)
		}
	}
	return returnStockExchange
}
