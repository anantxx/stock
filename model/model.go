package model

//jsonResponse Type
type jsonResponse struct {
	// Reserved field to add some meta information to the API response
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

//jsonErrorResponse Type
type jsonErrorResponse struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Error   string `json:"message"`
}

type Result struct {
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Price          string `json:"price"`
	CloseYesterday string `json:"close_yesterday"`
	Currency       string `json:"currency"`
	MarketCap      string `json:"market_cap"`
	Volume         string `json:"volume"`
	Timezone       string `json:"timezone"`
	TimezoneName   string `json:"timezone_name"`
	GmtOffset      string `json:"gmt_offset"`
	LastTradeTime  string `json:"last_trade_time"`
	StockExchange  string `json:"stock_exchange_short"`
}

type APIResponse struct {
	SymbolsRequested int                 `json:"symbols_requested"`
	SymbolsReturned  int                 `json:"symbols_returned"`
	Data             []map[string]string `json:"data"`
}
