package private

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/austinjhunt/go-gemini/public"
	"github.com/austinjhunt/go-gemini/util"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables
	util.Info("Initializing private.go")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Unable to load .env file:", err)
	}
}

func PostPrivateEndpoint(endpoint string, payload map[string]interface{}, target interface{}) error {
	/*
		Perform an HTTP POST request on a private Gemini API endpoint and unmarshal the JSON response into the provided target interface.

		Args:
		endpoint (string) - endpoint path (e.g., /v1/order/new) to make request against. Environment variable API_ENVIRONMENT determines whether to use api.gemini.com (prod) or api.sandbox.gemini.com (sandbox)
		target - any type of JSON object in which the JSON response gets stored, passed as &target (pointer to a variable in which response is to be stored) when method is invoked

		Returns nothing if successful, returns error if it fails
	*/
	payloadStr, err := json.Marshal(payload)
	if err != nil {
		return errors.New("Error encoding payload to string: " + err.Error())
	}
	util.Info("Posting payload to private endpoint " + endpoint + "\n\tPayload: " + string(payloadStr))
	// Ensure required API keys are set
	if os.Getenv("GEMINI_EXCHANGE_API_KEY") == "" || len([]byte(os.Getenv("GEMINI_EXCHANGE_API_SECRET"))) == 0 {
		log.Println("GEMINI_EXCHANGE_API_KEY and GEMINI_EXCHANGE_API_SECRET are not both present in environment")
	} else {
		log.Println("Private credentials from .env successfully loaded for private API functions!")
	}

	// Retrieve API key and secret from environment
	apiKey := util.GetEnvOrDefault("GEMINI_EXCHANGE_API_KEY", "")
	apiSecret := []byte(util.GetEnvOrDefault("GEMINI_EXCHANGE_API_SECRET", ""))

	// Build the full URL
	url := util.GetBaseAPIUrl() + endpoint

	// Create the nonce for the payload
	nonce := time.Now().UTC().Unix()

	// Merge the provided payload with auto-generated values
	mergedPayload := make(map[string]interface{})
	for key, value := range payload {
		mergedPayload[key] = value
	}
	mergedPayload["request"] = endpoint
	mergedPayload["nonce"] = strconv.FormatInt(nonce, 10)

	// Encode the payload to JSON
	payloadJSON, err := json.Marshal(mergedPayload)
	if err != nil {
		return errors.New("Error encoding payload: " + err.Error())
	}

	// Base64 encode the JSON payload
	b64Payload := base64.StdEncoding.EncodeToString(payloadJSON)

	// Create the HMAC signature using SHA384
	h := hmac.New(sha512.New384, apiSecret)
	h.Write([]byte(b64Payload))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	// Prepare the HTTP request headers
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return errors.New("Error creating request: " + err.Error())
	}

	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Content-Length", "0")
	req.Header.Set("X-GEMINI-APIKEY", apiKey)
	req.Header.Set("X-GEMINI-PAYLOAD", b64Payload)
	req.Header.Set("X-GEMINI-SIGNATURE", signature)
	req.Header.Set("Cache-Control", "no-cache")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("Error sending request: " + err.Error())
	}
	defer resp.Body.Close()

	// Read the response
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	// Handle non-OK HTTP status codes
	if resp.StatusCode != http.StatusOK {
		return errors.New("Error: status " + resp.Status + ", response: " + buf.String())
	}

	// Parse the response JSON into the target interface
	if err := json.Unmarshal(buf.Bytes(), target); err != nil {
		return errors.New("Error parsing response JSON: " + err.Error())
	}

	responseStr, err := json.Marshal(target)
	if err != nil {
		return errors.New("Error converting response to string: " + err.Error())
	}
	util.Info("Response from POST: \n\t" + string(responseStr))

	return nil
}

func GetClosedOrdersHistory() []Order {
	/*
		This API retrieves (closed) orders history for an account.

		The API key you use to access this endpoint must have the Trader or Auditor role assigned. See Roles for more information.
	*/

	util.Info("GetClosedOrdersHistory")

	var ordersHistory []Order
	// Create the payload map separately
	payload := map[string]interface{}{}

	endpoint := "/v1/orders/history"

	// Pass the payload to the function
	err := PostPrivateEndpoint(endpoint, payload, &ordersHistory)

	if err != nil {
		log.Fatalf("Error fetching orders history: %v", err)
		return nil
	}
	return ordersHistory
}

func GetOrderStatus(order_id string, opts *GetOrderStatusOptions) *Order {
	/*
		Get order status

		Required Args:
		order_id (string) - The order id to get information on. The order_id represents a whole number and is transmitted as an unsigned 64-bit integer in JSON format. An unsigned 64-bit integer is a non-negative whole number with a maximum value of 18,446,744,073,709,551,615. order_id cannot be used in combination with client_order_id.
		include_trades (bool) - if True the endpoint will return individual trade details of all fills from the order.

		Optional Args:
		account	(string) - Optional. Required for Master API keys as described in Private API Invocation. The name of the account within the subaccount group. Specifies the account on which the order was placed. Only available for exchange accounts.
		client_order_id	(string) - Optional. The client_order_id used when placing the order. client_order_id cannot be used in combination with order_id
		include_trades	(boolean) -	Optional. Either True or False. If True the endpoint will return individual trade details of all fills from the order

		Response: pointer to an Order object
	*/
	util.Info("GetOrderStatus")

	var orderStatus Order
	endpoint := "/v1/order/status"

	// Create the payload map separately
	payload := map[string]interface{}{
		"order_id": order_id,
	}

	// Pass the payload to the function
	err := PostPrivateEndpoint(endpoint, payload, &orderStatus)

	if err != nil {
		log.Fatalf("Error fetching order status: %v", err)
		return nil
	}
	return &orderStatus
}

func NewOrder(symbol string, amount string, price string, side string, orderType string, opts *NewOrderOptions) *Order {
	/*
		If you wish orders to be automatically cancelled when your session ends, see the require heartbeat section, or manually send the cancel all session orders message.

		Master API keys do not support cancelation on disconnect via heartbeat.

		Enabled for perpetuals accounts from July 10th, 0100hrs ET onwards.

		A Stop-Limit order is an order type that allows for order placement when a price reaches a specified level. Stop-Limit orders take in both a price and and a stop_price as parameters. The stop_price is the price that triggers the order to be placed on the continous live order book at the price. For buy orders, the stop_price must be below the price while sell orders require the stop_price to be greater than the price.

		What about market orders?
		The API doesn't directly support market orders because they provide you with no price protection.

		Instead, use the “immediate-or-cancel” order execution option, coupled with an aggressive limit price (i.e. very high for a buy order or very low for a sell order), to achieve the same result.

		Required Args:
		symbol (string) - symbol for the new order
		amount (string) - quoted decimal amount to purchase
		price (string) - quoted decimal amount to spend per unit
		side (string) - "buy" or "sell"
		type	(string) -	The order type. "exchange limit" for all order types except for stop-limit orders. "exchange stop limit" for stop-limit orders.

		Optional args:
		client_order_id (string) - (recommended) - client-specified order id
		options	(array) -	Optional. An optional array containing at most one supported order execution option. See Order execution options for details.
		stop_price	(string)	Optional. The price to trigger a stop-limit order. Only available for stop-limit orders.
		account	(string) -	Optional. Required for Master API keys as described in Private API Invocation. The name of the account within the subaccount group. Specifies the account on which you intend to place the order. Only available for exchange accounts.

		Response : map
		Response will be the same fields included in Order Status
	*/

	util.Info(fmt.Sprintf("NewOrder called with symbol: %s, amount: %s, price: %s, side: %s, orderType: %s, opts: %+v", symbol, amount, price, side, orderType, opts))

	var newOrder Order
	endpoint := "/v1/order/new"

	// Create the payload map separately
	payload := map[string]interface{}{
		"symbol": symbol,
		"amount": amount,
		"price":  price,
		"side":   side,
		"type":   orderType,
	}

	// Pass the payload to the function
	err := PostPrivateEndpoint(endpoint, payload, &newOrder)

	if err != nil {
		log.Fatalf("Error creating new order: %v", err)
		return nil
	}
	return &newOrder
}

func StopLimitBuy(symbol string, dollarAmount float64, stopPrice float64, limitPrice float64) *Order {
	/**
	  StopLimitBuy places a stop-limit buy order.

	  Parameters:
	  - symbol (string): The trading pair symbol (e.g., "BTCUSD").
	  - dollarAmount (float64): The amount in USD to spend on the buy order.
	  - stopPrice (float64): The price that triggers the order to be placed.
	  - limitPrice (float64): The price at which the order will be executed.

	  Returns:
	  - *Order: A pointer to the created order object.

	  Behavior:
	  - Validates that stopPrice is less than limitPrice.
	  - Calculates the amount to buy based on the dollarAmount and limitPrice.
	  - Uses the NewOrder function to place the stop-limit buy order.

	  Notes:
	  - The stopPrice must be less than the limitPrice for buy orders.
	  - Logs an error and exits if any validation fails or if fetching the current price fails.
	*/
	util.Info(fmt.Sprintf("StopLimitBuy called with symbol: %s, dollarAmount: %f, stopPrice: %f, limitPrice: %f", symbol, dollarAmount, stopPrice, limitPrice))

	// Fetch the current coin price
	currentPrice := public.GetCurrentCoinPriceUSD(symbol)
	if currentPrice < 0 {
		log.Fatalf("Failed to fetch current price for symbol: %s", symbol)
		return nil
	}

	// Validate the stop price and limit price
	if stopPrice >= limitPrice {
		log.Fatalf("Invalid stop and limit prices: stopPrice (%f) must be less than limitPrice (%f)", stopPrice, limitPrice)
		return nil
	}

	// Calculate the amount to buy
	amount := dollarAmount / limitPrice

	// Prepare optional parameters
	opts := &NewOrderOptions{
		StopPrice: strconv.FormatFloat(stopPrice, 'f', -1, 64),
	}

	// Place the stop-limit buy order
	return NewOrder(
		symbol,
		strconv.FormatFloat(amount, 'f', -1, 64),
		strconv.FormatFloat(limitPrice, 'f', -1, 64),
		"buy",
		"exchange stop limit",
		opts,
	)
}

func StopLimitSell(symbol string, amount float64, stopPrice float64, limitPrice float64) *Order {
	/**
	  StopLimitSell places a stop-limit sell order.

	  Parameters:
	  - symbol (string): The trading pair symbol (e.g., "BTCUSD").
	  - amount (float64): The amount of the asset to sell.
	  - stopPrice (float64): The price that triggers the order to be placed.
	  - limitPrice (float64): The price at which the order will be executed.

	  Returns:
	  - *Order: A pointer to the created order object.

	  Behavior:
	  - Validates that stopPrice is greater than limitPrice.
	  - Uses the NewOrder function to place the stop-limit sell order.

	  Notes:
	  - The stopPrice must be greater than the limitPrice for sell orders.
	  - Logs an error and exits if any validation fails or if fetching the current price fails.
	*/

	util.Info(fmt.Sprintf("StopLimitBuy called with symbol: %s, amount: %f, stopPrice: %f, limitPrice: %f", symbol, amount, stopPrice, limitPrice))

	// Fetch the current coin price
	currentPrice := public.GetCurrentCoinPriceUSD(symbol)
	if currentPrice < 0 {
		log.Fatalf("Failed to fetch current price for symbol: %s", symbol)
		return nil
	}

	// Validate the stop price and limit price
	if stopPrice <= limitPrice {
		log.Fatalf("Invalid stop and limit prices: stopPrice (%f) must be greater than limitPrice (%f)", stopPrice, limitPrice)
		return nil
	}

	// Prepare optional parameters
	opts := &NewOrderOptions{
		StopPrice: strconv.FormatFloat(stopPrice, 'f', -1, 64),
	}

	// Place the stop-limit sell order
	return NewOrder(
		symbol,
		strconv.FormatFloat(amount, 'f', -1, 64),
		strconv.FormatFloat(limitPrice, 'f', -1, 64),
		"sell",
		"exchange stop limit",
		opts,
	)
}

func CancelOrder(order_id int) *Order {
	util.Info(fmt.Sprintf("CancelOrder called with order_id: %v", order_id))
	var canceledOrder Order
	endpoint := "/v1/order/cancel"
	payload := map[string]interface{}{
		"order_id": order_id,
	}
	// Pass the payload to the function
	err := PostPrivateEndpoint(endpoint, payload, &canceledOrder)

	if err != nil {
		log.Fatalf("Error canceling order: %v", err)
		return nil
	}
	return &canceledOrder
}

func TradingBot(symbol string, tradingAmount float64) {
	util.Info(fmt.Sprintf("TradingBot called with symbol: %s, tradingAmount: %f", symbol, tradingAmount))

	for {
		// Fetch the current price
		currentPrice := public.GetCurrentCoinPriceUSD(symbol)
		if currentPrice == -1 {
			log.Printf("Failed to fetch current price for symbol: %s. Retrying...", symbol)
			continue
		}

		// Define the price range for trading
		buyStopPrice := 101.0
		buyLimitPrice := 101.5
		sellStopPrice := 109.0
		sellLimitPrice := 108.5

		if currentPrice <= buyStopPrice {
			// Place a stop-limit buy order
			log.Printf("Placing a buy order at stopPrice: %f, limitPrice: %f", buyStopPrice, buyLimitPrice)
			order := StopLimitBuy(symbol, tradingAmount, buyStopPrice, buyLimitPrice)
			if order != nil {
				log.Printf("Buy order placed: %+v", order)
			}
		} else if currentPrice >= sellStopPrice {
			// Calculate the amount to sell (assuming full balance)
			balance := tradingAmount / currentPrice // Simplified for example
			log.Printf("Placing a sell order at stopPrice: %f, limitPrice: %f", sellStopPrice, sellLimitPrice)
			order := StopLimitSell(symbol, balance, sellStopPrice, sellLimitPrice)
			if order != nil {
				log.Printf("Sell order placed: %+v", order)
			}
		}

		// Sleep for a while before checking again (e.g., to avoid hitting API rate limits)
		time.Sleep(10 * time.Second)
	}
}
