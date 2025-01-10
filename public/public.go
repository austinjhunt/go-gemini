package public

import (
	"log"
	"strconv"

	"github.com/austinjhunt/go-gemini/util"
)

// invoke Gemini exchange REST API public endpoints; one function per endpoint

func GetSymbols() []string {
	/* 
	Get a list of all available symbols

	Args:
	None

	Returns:
	The response will be an array of strings
	*/
	var symbols []string
	url := "/v1/symbols"

	err := util.GetPublicEndpoint(url, &symbols)
	if err != nil {
		log.Fatalf("Error fetching symbols: %v", err)
	}

	return symbols
}

func GetSymbolDetails(symbol string) map[string]interface{} {
	/*
	Get extra detail on supported symbols, such as minimum order size, tick size, quote increment and more

	Args:
	symbol (string): Trading pair symbol. See symbols and minimums
	
	*/

	var details map[string]interface{}
	url := "/v1/symbols/details/" + symbol

	err := util.GetPublicEndpoint(url, &details)
	if err != nil {
		log.Fatalf("Error fetching symbol details: %v", err)
	}

	return details
}

func GetNetwork(token string) map[string]interface{} {
	var network map[string]interface{}
	url := "/v1/network/" + token

	err := util.GetPublicEndpoint(url, &network)
	if err != nil {
		log.Fatalf("Error fetching network: %v", err)
	}

	return network
}

func GetTicker(symbol string) map[string]interface{} {
	var ticker map[string]interface{}
	url := "/v1/pubticker/" + symbol

	err := util.GetPublicEndpoint(url, &ticker)
	if err != nil {
		log.Fatalf("Error fetching ticker: %v", err)
	}

	return ticker
}

func GetTickerV2(symbol string) map[string]interface{} {
	var ticker map[string]interface{}
	url := "https://api.gemini.com/v2/ticker/" + symbol

	err := util.GetPublicEndpoint(url, &ticker)
	if err != nil {
		log.Fatalf("Error fetching ticker v2: %v", err)
	}

	return ticker
}

func GetCandles(symbol string, time_frame string) [][]interface{} {
	var candles [][]interface{}
	url := "https://api.gemini.com/v2/candles/" + symbol + "/" + time_frame

	err := util.GetPublicEndpoint(url, &candles)
	if err != nil {
		log.Fatalf("Error fetching candles: %v", err)
	}

	return candles
}


func GetDerivativesCandles(symbol string, time_frame string) [][]interface{} {
	/* 
	Get time-intervaled data for the provided perps symbol

	Args: 
	symbol (string): Trading pair symbol. Available only for perpetual pairs like BTCGUSDPERP, See symbols and minimums
	time_frame (string): Time range for each candle. 1m: 1 minute (only)

	Returns:
	The response will be an array of arrays
	*/
	var derivativesCandles [][]interface{}
	 
	url := "https://api.gemini.com/v2/derivatives/candles/" + symbol + "/" + time_frame

	err := util.GetPublicEndpoint(url, &derivativesCandles)

	if err != nil {
		log.Fatalf("Error fetching Derivatives Candles %v", err)
	}
	return derivativesCandles
}

func GetFeePromos() map[string]interface{} {
	/* 
	Get symbols that currently have fee promos

	Args: 
	None

	Returns:
	The response will be a JSON object
	*/
	 
	var feePromos map[string]interface{}

	url := "/v1/feepromos"
	err := util.GetPublicEndpoint(url, &feePromos) 
	if err != nil {
		log.Fatalf("Error fetching fee promos %v", err)
	}
	return feePromos
}

func GetCurrentOrderBook(symbol string) map[string]interface{} {
	/* 
	Return the current order book as two arrays (bids / asks)

	Args: 
	symbol (string): Trading pair symbol. See symbols and minimums

	Returns:
	The response will be two arrays
	*/
	 
	var currentOrderBook map[string]interface{}

	url := "/v1/book/" + symbol

	err := util.GetPublicEndpoint(url, &currentOrderBook)

	if err != nil {
		log.Fatalf("Error fetching Current Order Book %v", err)
	}
	
	return currentOrderBook
}

func GetTradeHistory(symbol string) []map[string]interface{} {
	/* 
	Return the trades that have executed since the specified timestamp

	Args: 
	symbol (string): Trading pair symbol. See symbols and minimums

	Returns:
	The response will be an array of JSON objects, sorted by timestamp, with the newest trade shown first
	*/
	var tradeHistory []map[string]interface{}
	 
	url := "/v1/trades/" + symbol
	err := util.GetPublicEndpoint(url, &tradeHistory)
	
	if err != nil {
		log.Fatalf("Error fetching Trade History %v", err)
	}
	return tradeHistory
}

func GetPriceFeed() []map[string]interface{} {
	/* 
	Return a list of objects, one for each pair, with the current price and 24 hour change in price

	Args: 
	None

	Returns:
	Response is a list of objects, one for each pair, with the following fields
	*/
	var priceFeed []map[string]interface{}
	
	url := "/v1/pricefeed"
	err := util.GetPublicEndpoint(url, &priceFeed)
	if err != nil {
		log.Fatalf("Error fetching price feed %v", err)
	}
	return priceFeed
}

func GetFundingAmount(symbol string) map[string]interface{} {
	/* 
	Get extra detail on supported symbols, such as minimum order size, tick size, quote increment and more

	Args: 
	symbol (string): Trading pair symbol. See symbols and minimums

	Returns:
	The response will be an object
	*/
	var fundingAmount map[string]interface{}
	url := "/v1/fundingamount/" + symbol

	err := util.GetPublicEndpoint(url, &fundingAmount) 
	if err != nil {
		log.Fatalf("Error fetching funding amount %v", err)
	}
	return fundingAmount
}
func DownloadFundingAmountReport(symbol string, fromDate string, toDate string, numRows int) error {
	/*
	Downloads a CSV or Excel file with funding amount records.

	Args: 
		symbol (string): Trading pair symbol. See symbols and minimums.
		fromDate (string): Start date for the report (format: yyyy-MM-dd).
		toDate (string): End date for the report (format: yyyy-MM-dd).
		numRows (int): Maximum number of records to fetch.

	Returns:
		error: Returns an error if the file cannot be downloaded.
	*/

	// Construct the URL for the report.
	url := "/v1/fundingamountreport/records.xlsx?symbol=" + symbol + "&fromDate=" + fromDate + "&toDate=" + toDate + "&numRows=" + strconv.Itoa(numRows)
	
	// Define the file path to save the downloaded file.
	filePath := "funding_amount_report_" + symbol + "_" + fromDate + "_to_" + toDate + ".xlsx"

	err := util.DownloadPublicFile(url, filePath)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	return nil
}

