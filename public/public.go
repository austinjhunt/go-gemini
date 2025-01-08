package public 

import (
	"gogemini/util"
)

// invoke Gemini exchange REST API public endpoints; one function per endpoint

/*
Use function name pattern <Method><Endpoint> for each endpoint e.g., GetSymbols, GetSymbolDetails, GetNetwork, GetTicker, GetTickerV2, GetCandles, GetDerivativesCandles, GetFeePromos, GetCurrentOrderBook, GetTradeHistory, GetPriceFeed, GetFundingAmount, GetFundingAmountReportFile
*/
 

// GetSymbols - return JSON data, do not just print
func GetSymbols() []byte {
	url := "https://api.gemini.com/v1/symbols"
	return util.GetPublicEndpoint(url)
}

