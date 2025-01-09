 
package public 

import (
	"testing"
	"log"
	"github.com/austinjhunt/go-gemini/util"
)

/* valid responses in API reference to use for testing 


/*

Full API reference:
Symbols
import requests, json

base_url = "https://api.gemini.com/v1"
response = requests.get(base_url + "/symbols")
symbols = response.json()

print(symbols)
This endpoint retrieves all available symbols for trading

HTTP Request
GET https://api.gemini.com/v1/symbols

Sandbox
Base URL can be changed to api.sandbox.gemini.com/v1 for test purposes.

URL Parameters
None

Response
An array of supported symbols. The full list of supported symbols:

Symbols
btcusd ethbtc ethusd bchusd bchbtc bcheth ltcusd ltcbtc ltceth ltcbch batusd daiusd linkusd oxtusd linkbtc linketh ampusd compusd paxgusd mkrusd zrxusd manausd storjusd crvusd uniusd renusd umausd yfiusd aaveusd filusd btceur btcgbp etheur ethgbp btcsgd ethsgd sklusd grtusd lrcusd sandusd cubeusd lptusd maticusd injusd sushiusd dogeusd ftmusd ankrusd btcgusd ethgusd ctxusd xtzusd axsusd efilfil gusdusd dogebtc dogeeth rareusd qntusd maskusd fetusd api3usd usdcusd shibusd rndrusd galausd ensusd zecusd ldousd solusd apeusd gusdsgd zbcusd chzusd jamusd gmtusd aliusd gusdgbp dotusd ernusd galusd samousd imxusd iotxusd avaxusd atomusd usdtusd btcusdt ethusdt pepeusd xrpusd hntusd btcgusdperp ethgusdperp pepegusdperp xprgusdperp solgusdperp maticgusdperp dogegusdperp linkgusdperp avaxgusdperp ltcgusdperp dotgusdperp bnbgusdperp injgusdperp wifgusdperp wifusd bonkusd popcatusd opusd bonkgusdperp popcatgusdperp opgusdperp moodengusd pnutusd goatusd mewusd bomeusd flokiusd pythusd solbtc soleth bchgusdperp bomegusdperp bonkgusdperp flokigusdperp goatgusdperp mewgusdperp moodenggusdperp pnutgusdperp polgusdperp pythgusdperp shibgusdperp unigusdperp chillguyusd

JSON response


["aaveusd","aliusd","ampusd","ankrusd","apeusd","api3usd","atomusd","avaxgusdperp","avaxusd","axsusd","batusd","bchbtc","bcheth","bchgusdperp","bchusd","bnbgusdperp","bomegusdperp","bomeusd","bonkgusdperp","bonkusd","btceur","btcgbp","btcgusd","btcgusdperp","btcsgd","btcusd","btcusdt","chillguyusd","chzusd","compusd","crvusd","ctxusd","cubeusd","daiusd","dogebtc","dogeeth","dogegusdperp","dogeusd","dotgusdperp","dotusd","efilfil","elonusd","ensusd","ernusd","ethbtc","etheur","ethgbp","ethgusd","ethgusdperp","ethsgd","ethusd","ethusdt","fetusd","filusd","flokigusdperp","flokiusd","ftmusd","galausd","galusd","gmtusd","goatgusdperp","goatusd","grtusd","gusdgbp","gusdsgd","gusdusd","hntusd","hypegusdperp","imxusd","injgusdperp","injusd","iotxusd","jamusd","ldousd","linkbtc","linketh","linkgusdperp","linkusd","lptusd","lrcusd","ltcbch","ltcbtc","ltceth","ltcgusdperp","ltcusd","manausd","maskusd","maticgusdperp","maticusd","mewgusdperp","mewusd","mkrusd","moodenggusdperp","moodengusd","opgusdperp","opusd","oxtusd","paxgusd","pepegusdperp","pepeusd","pnutgusdperp","pnutusd","polgusdperp","popcatgusdperp","popcatusd","pythgusdperp","pythusd","qntusd","rareusd","renusd","rndrusd","samousd","sandusd","shibgusdperp","shibusd","sklusd","solbtc","soleth","solgusdperp","solusd","storjusd","sushiusd","umausd","unigusdperp","uniusd","usdcusd","usdtgusd","usdtusd","wifgusdperp","wifusd","xrpgusdperp","xrpusd","xtzusd","yfiusd","zecusd","zrxusd"]

Symbol Details
import requests, json

base_url = "https://api.gemini.com/v1"
response = requests.get(base_url + "/symbols/details/:symbol")
symbols = response.json()

print(symbols)
This endpoint retrieves extra detail on supported symbols, such as minimum order size, tick size, quote increment and more.

HTTP Request
GET https://api.gemini.com/v1/symbols/details/:symbol

Sandbox
Base URL can be changed to api.sandbox.gemini.com/v1 for test purposes.

URL Parameters
Parameter	Type	Description	Values
:symbol	String	Trading pair symbol	BTCUSD, etc. See symbols and minimums
Response
The response will be an object

Field	Type	Description
symbol	string	The requested symbol. See symbols and minimums
base_currency	string	CCY1 or the top currency. (ieBTC in BTCUSD)
quote_currency	string	CCY2 or the quote currency. (ie USD in BTCUSD)
tick_size	decimal	The number of decimal places in the base_currency. (ie 1e-8)
quote_increment	decimal	The number of decimal places in the quote_currency (ie 0.01)
min_order_size	string	The minimum order size in base_currency units (ie 0.00001)
status	string	Status of the current order book. Can be open, closed, cancel_only, post_only, limit_only.
wrap_enabled	boolean	When True, symbol can be wrapped using this endpoint: POST https://api.gemini.com/v1/wrap/:symbol
product_type	string	instrument type spot / swap -- where swap signifies perpetual swap.
contract_type	string	vanilla / linear / inverse where vanilla is for spot
while linear is for perpetual swap
and inverse is a special case perpetual swap where the perpetual contract will be settled in base currency
contract_price_currency	string	CCY2 or the quote currency for spot instrument (i.e. USD in BTCUSD)
Or collateral currency of the contract in case of perpetual swap instrument
JSON response - Spot instrument response

{
  "symbol": "BTCUSD",
  "base_currency": "BTC",
  "quote_currency": "USD",
  "tick_size": 1E-8,
  "quote_increment": 0.01,
  "min_order_size": "0.00001",
  "status": "open",
  "wrap_enabled": false,
  "product_type":"spot",
  "contract_type":"vanilla",
  "contract_price_currency":"USD"

}
JSON response - Perpetual Swap instrument response

{
  "symbol":"BTCETHPERP",
  "base_currency":"BTC",
  "quote_currency":"ETH",
  "tick_size":0.0001,
  "quote_increment":0.5,
  "min_order_size":"0.0001",
  "status":"open",
  "wrap_enabled":false,
  "product_type":"swap",
  "contract_type":"linear",
  "contract_price_currency":"GUSD"
}
Network
import requests, json

base_url = "https://api.gemini.com/v1"
response = requests.get(base_url + "/nework/rbn")
network = response.json()

print(network)
This endpoint retrieves the associated network for a requested token.

HTTP Request
GET https://api.gemini.com/v1/network/:token

Sandbox
Base URL can be changed to api.sandbox.gemini.com/v1 for test purposes.

URL Parameters
Parameter	Type	Description	Values
:token	String	Token identifier	BTC, ETH, SOL etc. See symbols and minimums
Response
The response will be a JSON object:

Field	Type	Description
token	string	The requested token.
network	array	Network of the requested token.
JSON response


{
  "token":"RBN",
  "network":["ethereum"]
}

Ticker
import requests, json

base_url = "https://api.gemini.com/v1"
response = requests.get(base_url + "/pubticker/btcusd")
btc_data = response.json()

print(btc_data)
JSON response

{
    "ask": "977.59",
    "bid": "977.35",
    "last": "977.65",
    "volume": {
        "BTC": "2210.505328803",
        "USD": "2135477.463379586263",
        "timestamp": 1483018200000
    }
}
This endpoint retrieves information about recent trading activity for the symbol.

HTTP Request
GET https://api.gemini.com/v1/pubticker/:symbol

URL Parameters
None

Sandbox
Base URL can be changed to api.sandbox.gemini.com/v1 for test purposes.

Response
The response will be an object

Field	Type	Description
bid	decimal	The highest bid currently available
ask	decimal	The lowest ask currently available
last	decimal	The price of the last executed trade
volume	node (nested)	Information about the 24 hour volume on the exchange. See below
The volume field will contain information about the 24 hour volume on the exchange. The volume is updated every five minutes based on a trailing 24-hour window of volume. It will have three fields

Field	Type	Description
timestamp	timestampms	The end of the 24-hour period over which volume was measured
(price symbol, e.g. "USD")	decimal	The volume denominated in the price currency
(quantity symbol, e.g. "BTC")	decimal	The volume denominated in the quantity currency
Ticker V2
import requests, json

base_url = "https://api.gemini.com/v2"
response = requests.get(base_url + "/ticker/btcusd")
btc_data = response.json()

print(btc_data)
JSON response

{
  "symbol": "BTCUSD",
  "open": "9121.76",
  "high": "9440.66",
  "low": "9106.51",
  "close": "9347.66",
  "changes": [
    "9365.1",
    "9386.16",
    "9373.41",
    "9322.56",
    "9268.89",
    "9265.38",
    "9245",
    "9231.43",
    "9235.88",
    "9265.8",
    "9295.18",
    "9295.47",
    "9310.82",
    "9335.38",
    "9344.03",
    "9261.09",
    "9265.18",
    "9282.65",
    "9260.01",
    "9225",
    "9159.5",
    "9150.81",
    "9118.6",
    "9148.01"
  ],
  "bid": "9345.70",
  "ask": "9347.67"
}
This endpoint retrieves information about recent trading activity for the provided symbol.

HTTP Request
GET https://api.gemini.com/v2/ticker/:symbol

URL Parameters
None

Sandbox
Base url can be changed to api.sandbox.gemini.com/v2 for test purposes.

Response
The response will be an object

Field	Type	Description
symbol	string	BTCUSD etc.
open	decimal	Open price from 24 hours ago
high	decimal	High price from 24 hours ago
low	decimal	Low price from 24 hours ago
close	decimal	Close price (most recent trade)
changes	array of decimals	Hourly prices descending for past 24 hours
--	decimal	Close price for each hour
bid	decimal	Current best bid
ask	decimal	Current best offer
Candles
import requests, json

base_url = "https://api.gemini.com/v2"
response = requests.get(base_url + "/candles/btcusd/15m")
btc_candle_data = response.json()

print(btc_candle_data)
JSON response

[
    [
     1559755800000,
     7781.6,
     7820.23,
     7776.56,
     7819.39,
     34.7624802159
    ],
    [1559755800000,
    7781.6,
    7829.46,
    7776.56,
    7817.28,
    43.4228281059],
    ...
]
This endpoint retrieves time-intervaled data for the provided symbol.

HTTP Request
GET https://api.gemini.com/v2/candles/:symbol/:time_frame

URL Parameters
None

Parameter	Type	Description	Values
:symbol	String	Trading pair symbol	BTCUSD, etc. See symbols and minimums
:time_frame	String	Time range for each candle	1m: 1 minute
5m: 5 minutes
15m: 15 minutes
30m: 30 minutes
1hr: 1 hour
6hr: 6 hours
1day: 1 day
Sandbox
Base URL can be changed to api.sandbox.gemini.com/v2 for test purposes.

Response
The response will be an array of arrays

Field	Type	Description
Array of Arrays	Descending order by time
-- -- time	long	Time in milliseconds
-- -- open	decimal	Open price
-- -- high	decimal	High price
-- -- low	decimal	Low price
-- -- close	decimal	Close price
-- -- volume	decimal	Volume
Derivatives Candles
import requests, json

base_url = "https://api.gemini.com/v2"
response = requests.get(base_url + "/derivatives/candles/BTCGUSDPERP/1m")
btc_perps_candle_data = response.json()

print(btc_perps_candle_data)
JSON response

[
    [
      1714126740000,
      68038,
      68038,
      68038,
      68038,
      0
    ],
    [
      1714126680000,
      68038,
      68038,
      68038,
      68038,
      0
    ],
    ...
]
This endpoint retrieves time-intervaled data for the provided perps symbol.

HTTP Request
GET https://api.gemini.com/v2/derivatives/candles/:symbol/:time_frame

URL Parameters
None

Parameter	Type	Description	Values
:symbol	String	Trading pair symbol	available only for perpetual pairs likeBTCGUSDPERP, See symbols and minimums
:time_frame	String	Time range for each candle	1m: 1 minute (only)
Sandbox
Base URL can be changed to api.sandbox.gemini.com/v2 for test purposes.

Response
The response will be an array of arrays

Field	Type	Description
Array of Arrays	Descending order by time
-- -- time	long	Time in milliseconds
-- -- open	decimal	Open price
-- -- high	decimal	High price
-- -- low	decimal	Low price
-- -- close	decimal	Close price
-- -- volume	decimal	Volume
Fee Promos
import requests, json

base_url = "https://api.gemini.com/v1"
response = requests.get(base_url + "/feepromos")
feepromos = response.json()

print(feepromos)
This endpoint retrieves symbols that currently have fee promos.

HTTP Request
GET https://api.gemini.com/v1/feepromos

Sandbox
Base URL can be changed to api.sandbox.gemini.com/v1 for test purposes.

URL Parameters
None

Response
The response will be a JSON object:

Field	Type	Description
symbols	array	Symbols that currently have fee promos
JSON response


{
  "symbols": [
    "GMTUSD",
    "GUSDGBP",
    "MIMUSD",
    "ORCAUSD",
    "RAYUSD",
    "FIDAUSD",
    "SOLUSD",
    "USDCUSD",
    "SRMUSD",
    "SBRUSD",
    "GUSDSGD",
    "DAIUSD"
  ]
}

Current Order Book
import requests, json

base_url = "https://api.gemini.com/v1"
response = requests.get(base_url + "/book/btcusd")
btc_book = response.json()

print(btc_book)
JSON response


{
  "bids": [{
            "price": "3607.85",
            "amount": "6.643373",
            "timestamp": "1547147541"
           }
           ...
           ],
  "asks": [{
            "price": "3607.86",
            "amount": "14.68205084",
            "timestamp": "1547147541"
           }
           ...
           ]
}
This will return the current order book as two arrays (bids / asks).

The quantities and prices returned are returned as strings rather than numbers. The numbers returned are exact, not rounded, and it can be dangerous to treat them as floating point numbers.
HTTP Request
GET https://api.gemini.com/v1/book/:symbol

Sandbox
Base URL can be changed to api.sandbox.gemini.com/v1 for test purposes.

URL Parameters
If a limit is specified on a side, then the orders closest to the midpoint of the book will be the ones returned.

Parameter	Type	Description
limit_bids	integer	Optional. Limit the number of bid (offers to buy) price levels returned. Default is 50. May be 0 to return the full order book on this side.
limit_asks	integer	Optional. Limit the number of ask (offers to sell) price levels returned. Default is 50. May be 0 to return the full order book on this side.
Response
The response will be two arrays

Field	Type	Description
bids	array	The bid price levels currently on the book. These are offers to buy at a given price
asks	array	The ask price levels currently on the book. These are offers to sell at a given price
The bids and the asks are grouped by price, so each entry may represent multiple orders at that price. Each element of the array will be a JSON object

Field	Type	Description
price	decimal	The price
amount	decimal	The total quantity remaining at the price
timestamp	timestamp	DO NOT USE - this field is included for compatibility reasons only and is just populated with a dummy value.
Trade History
import requests, json

base_url = "https://api.gemini.com/v1"
response = requests.get(base_url + "/trades/btcusd")
btcusd_trades = response.json()

print(btcusd_trades)
JSON response

[
  {
    "timestamp": 1547146811,
    "timestampms": 1547146811357,
    "tid": 5335307668,
    "price": "3610.85",
    "amount": "0.27413495",
    "exchange": "gemini",
    "type": "buy"
  },
  ...
]
This public API endpoint is limited to retrieving seven calendar days of data.

Please contact us through this form for information about Gemini market data.
This will return the trades that have executed since the specified timestamp. Timestamps are either seconds or milliseconds since the epoch (1970-01-01). See the Data Types section about timestamp for information on this.

Each request will show at most 500 records.

If no since or timestamp is specified, then it will show the most recent trades; otherwise, it will show the most recent trades that occurred after that timestamp.

HTTP Request
GET https://api.gemini.com/v1/trades/:symbol

Sandbox
Base URL can be changed to api.sandbox.gemini.com/v1 for test purposes.

URL Parameters
Parameter	Type	Description
timestamp	timestamp	Optional. Only return trades after this timestamp. See Data Types: Timestamps for more information. If not present, will show the most recent trades. For backwards compatibility, you may also use the alias 'since'. With timestamp, there is a 90-day hard limit.
since_tid	integer	Optional. Only retuns trades that executed after this tid. since_tid trumps timestamp parameter which has no effect if provided too. You may set since_tid to zero to get the earliest available trade history data.
limit_trades	integer	Optional. The maximum number of trades to return. The default is 50.
include_breaks	boolean	Optional. Whether to display broken trades. False by default. Can be '1' or 'true' to activate
Response
The response will be an array of JSON objects, sorted by timestamp, with the newest trade shown first.

Field	Type	Description
timestamp	timestamp	The time that the trade was executed
timestampms	timestampms	The time that the trade was executed in milliseconds
tid	integer	The trade ID number
price	decimal	The price the trade was executed at
amount	decimal	The amount that was traded
exchange	string	Will always be "gemini"
type	string	
buy means that an ask was removed from the book by an incoming buy order
sell means that a bid was removed from the book by an incoming sell order
broken	boolean	Whether the trade was broken or not. Broken trades will not be displayed by default; use the include_breaks to display them.
Price feed
import requests, json

base_url = "https://api.gemini.com/v1"
response = requests.get(base_url + "/pricefeed")
prices = response.json()

print(prices)
JSON Response

[   
    {
        "pair":"BTCUSD",
        "price":"9500.00",
        "percentChange24h": "5.23"
    },
    {
        "pair":"ETHUSD",
        "price":"257.54",
        "percentChange24h": "4.85"
    },
    {
        "pair":"BCHUSD",
        "price":"450.10",
        "percentChange24h": "-2.91"
    },
    {
        "pair":"LTCUSD",
        "price":"79.50",
        "percentChange24h": "7.63"
    }
]
HTTP Request
GET https://api.gemini.com/v1/pricefeed

Sanbox
Base URL can be changed to api.sandbox.gemini.com/v1 for test purposes.

URL Parameters
None

Response
Response is a list of objects, one for each pair,with the following fields:

Field	Type	Description
pair	String	Trading pair symbol. See symbols and minimums
price	String	Current price of the pair on the Gemini order book
percentChange24h	String	24 hour change in price of the pair on the Gemini order book
Funding Amount
import requests, json

base_url = "https://api.gemini.com/v1"
response = requests.get(base_url + "/fundingamount/:symbol")
symbols = response.json()

print(symbols)
This endpoint retrieves extra detail on supported symbols, such as minimum order size, tick size, quote increment and more.

HTTP Request
GET https://api.gemini.com/v1/fundingamount/:symbol

Sandbox
Base URL can be changed to api.sandbox.gemini.com/v1 for test purposes.

URL Parameters
Parameter	Type	Description	Values
:symbol	String	Trading pair symbol	BTCGUSDPERP, etc. See symbols and minimums
Response
The response will be an object

Field	Type	Description
symbol	string	The requested symbol. See symbols and minimums
fundingDateTime	string	UTC date time in format yyyy-MM-ddThh:mm:ss.SSSZ format
fundingTimestampMilliSecs	long	Current funding amount Epoc time.
nextFundingTimestamp	long	Next funding amount Epoc time.
amount	decimal	The dollar amount for a Long 1 position held in the symbol for funding period (1 hour)
estimatedFundingAmount	decimal	The estimated dollar amount for a Long 1 position held in the symbol for next funding period (1 hour)
JSON response


{
    "symbol": "btcgusdperp",
    "fundingDateTime": "2023-06-12T03:00:00.000Z",
    "fundingTimestampMilliSecs": 1686538800000,
    "nextFundingTimestamp": 1686542400000,
    "fundingAmount": 0.51692,
    "estimatedFundingAmount": 0.27694
}

Funding Amount Report File
HTTP Request
GET https://api.gemini.com/v1/fundingamountreport/records.xlsx?symbol=<symbol>&fromDate=<date>&toDate=<date>&numRows=<rows>

Parameters
Parameter	Type	Description
symbol	string	Mandatory
fromDate	integer	Mandatory if toDate is specified. Else, Optional. If empty, will only fetch records by numRows value
toDate	string	Mandatory if fromDate is specified. Else, Optional. If empty, will only fetch records by numRows value
numRows	integer	Optional. If empty, default value '8760'
Examples
symbol=BTCGUSDPERP&fromDate=2024-04-10&toDate=2024-04-25&numRows=1000
Compare and obtain the minimum records between (2024-04-10 to 2024-04-25) and 1000. If (2024-04-10 to 2024-04-25) contains 360 records. Then fetch the minimum between 360 and 1000 records only.

symbol=BTCGUSDPERP&numRows=2024-04-10&toDate=2024-04-25
If (2024-04-10 to 2024-04-25) contains 360 records. Then fetch 360 records only.

symbol=BTCGUSDPERP&numRows=1000
Fetch maximum 1000 records starting from Now to a historical date

symbol=BTCGUSDPERP
Fetch maximum 8760 records starting from Now to a historical date

Response
csv / excel file will be downloaded

 
*/


// TestGetSymbols tests the GetSymbols function
func TestGetSymbols(t *testing.T) {
	response := GetSymbols()
	log.Println(response)
	// response should be an array of supported symbols
	if len(response) == 0 {
		t.Errorf("GetSymbols failed")
	}
	// example symbols that should be in the response: btcusd, ethusd, ethbtc, bchusd, bchbtc, bcheth, ltcusd, ltcbtc, ltceth, ltcbch, batusd, daiusd, linkusd, oxtusd, linkbtc, linketh, ampusd, compusd, paxgus
	expectedSymbols := []string{"btcusd", "ethusd", "ethbtc", "bchusd", "bchbtc", "bcheth", "ltcusd", "ltcbtc", "ltceth", "ltcbch", "batusd", "daiusd", "linkusd", "oxtusd", "linkbtc", "linketh", "ampusd", "compusd", "paxgusd"}
	// if response does not contain all expected symbols, test failed
	for _, symbol := range expectedSymbols {
		// if response does not contain the symbol, test failed
		if !util.ArrayContainsString(response, symbol) { 
			t.Errorf("GetSymbols failed")
		}
	} 
}

// TestGetSymbolDetails tests the GetSymbolDetails function
func TestGetSymbolDetails(t *testing.T) {
	response := GetSymbolDetails("btcusd")
	log.Println(response)
	if len(response) == 0 {
		t.Errorf("GetSymbolDetails failed")
	}
}

// TestGetNetwork tests the GetNetwork function
func TestGetNetwork(t *testing.T) {
	response := GetNetwork("btc")
	log.Println(response)
	if len(response) == 0 {
		t.Errorf("GetNetwork failed")
	}
}

// TestGetTicker tests the GetTicker function
func TestGetTicker(t *testing.T) {
	response := GetTicker("btcusd")
	log.Println(response)
	if len(response) == 0 {
		t.Errorf("GetTicker failed")
	}
}

// TestGetTickerV2 tests the GetTickerV2 function
func TestGetTickerV2(t *testing.T) {
	response := GetTickerV2("btcusd")
	log.Println(response)
	if len(response) == 0 {
		t.Errorf("GetTickerV2 failed")
	}
}

// TestGetCandles tests the GetCandles function
func TestGetCandles(t *testing.T) {
	response := GetCandles("btcusd", "1m")
	log.Println(response)
	if len(response) == 0 {
		t.Errorf("GetCandles failed")
	}
}
 