package public

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/austinjhunt/go-gemini/util"
)

// invoke Gemini exchange REST API public endpoints; one function per endpoint

func GetPublicEndpoint(endpoint string, target interface{}) error {
	/*
		Perform an HTTP GET request on a public Gemini API endpoint and unmarshal the JSON response into the provided target interface.

		Args:
		endpoint (string) - API endpoint to use
		target - any type of JSON object in which the JSON response gets stored, passed as &target (pointer to a variable in which response is to be stored) when method is invoked

		Returns nothing if successful, returns error if it fails

	*/

	url := util.GetBaseAPIUrl() + endpoint

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return errors.New("error creating request: " + err.Error())
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return errors.New("error making request: " + err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("received non-200 status code: " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.New("error reading response body: " + err.Error())
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return errors.New("error unmarshalling response: " + err.Error())
	}

	return nil
}

func DownloadPublicFile(endpoint string, filePath string) error {
	/*
		Download a file from a url to a specific file path

		Args:
			endpoint (string): endpoint from which to download file
			filePath (string): path on local file system to which downloaded file will be saved

		Returns:
			error: Returns an error if the file cannot be downloaded.
	*/

	// Perform the HTTP GET request to download the file.
	url := util.GetBaseAPIUrl() + endpoint
	response, err := http.Get(url)
	if err != nil {
		return errors.New("failed to download funding amount report: " + err.Error())
	}
	defer response.Body.Close()

	// Check if the response status is OK.
	if response.StatusCode != http.StatusOK {
		return errors.New("unexpected HTTP status: " + response.Status)
	}

	// Create the file locally.
	file, err := os.Create(filePath)
	if err != nil {
		return errors.New("failed to create the file: " + err.Error())
	}
	defer file.Close()

	// Copy the response body to the file.
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return errors.New("failed to save the report: " + err.Error())
	}

	log.Printf("Report downloaded successfully and saved to %s\n", filePath)
	return nil
}

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

	err := GetPublicEndpoint(url, &symbols)
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

	err := GetPublicEndpoint(url, &details)
	if err != nil {
		log.Fatalf("Error fetching symbol details: %v", err)
	}

	return details
}

func GetNetwork(token string) map[string]interface{} {
	var network map[string]interface{}
	url := "/v1/network/" + token

	err := GetPublicEndpoint(url, &network)
	if err != nil {
		log.Fatalf("Error fetching network: %v", err)
	}

	return network
}

func GetTicker(symbol string) *TickerV1 {
	util.Info(fmt.Sprintf("GetTicker, symbol: %s", symbol))
	var ticker TickerV1
	url := "/v1/pubticker/" + symbol

	err := GetPublicEndpoint(url, &ticker)
	if err != nil {
		log.Fatalf("Error fetching ticker: %v", err)
		return nil
	}

	return &ticker
}

func GetTickerV2(symbol string) *TickerV2 {
	util.Info(fmt.Sprintf("GetTickerV2, symbol: %s", symbol))
	var ticker TickerV2
	url := "/v2/ticker/" + strings.ToLower(symbol)

	err := GetPublicEndpoint(url, &ticker)
	if err != nil {
		log.Fatalf("Error fetching ticker v2: %v", err)
		return nil
	}

	return &ticker
}

func GetCandles(symbol string, time_frame string) [][]interface{} {
	var candles [][]interface{}
	url := "/v2/candles/" + symbol + "/" + time_frame

	err := GetPublicEndpoint(url, &candles)
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

	url := "/v2/derivatives/candles/" + symbol + "/" + time_frame

	err := GetPublicEndpoint(url, &derivativesCandles)

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
	err := GetPublicEndpoint(url, &feePromos)
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

	err := GetPublicEndpoint(url, &currentOrderBook)

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
	err := GetPublicEndpoint(url, &tradeHistory)

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
	err := GetPublicEndpoint(url, &priceFeed)
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

	err := GetPublicEndpoint(url, &fundingAmount)
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

	err := DownloadPublicFile(url, filePath)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	return nil
}

func GetCurrentCoinPriceUSD(symbol string) float64 {
	ticker := GetTickerV2(symbol)
	if ticker == nil {
		log.Fatalf("Ticker data for %s not found", symbol)
		return -1
	}

	askPrice, err := strconv.ParseFloat(ticker.Ask, 64)
	if err != nil {
		log.Fatalf("Invalid ask price for %s: %v", symbol, askPrice)
		return -1
	}

	return askPrice
}

func ConvertUSDToCryptoAmount(dollarAmount float64, symbol string) float64 {
	util.Info(fmt.Sprintf("Converting %f USD to %s", dollarAmount, symbol))
	ticker := GetTickerV2(symbol)
	if ticker == nil {
		log.Fatalf("Ticker data for %s not found", symbol)
		return -1
	}

	askPrice, err := strconv.ParseFloat(ticker.Ask, 64)
	if err != nil {
		log.Fatalf("Invalid ask price for %s: %v", symbol, askPrice)
		return -1
	}

	cryptoAmount := dollarAmount / askPrice
	util.Info(fmt.Sprintf("%f USD = %f %s", dollarAmount, cryptoAmount, symbol))
	return cryptoAmount
}
