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

func PostPrivateEndpoint(payload []byte, target interface{}) error {
	/*
				Perform an HTTP POST request on a private Gemini API endpoint and unmarshal the JSON response into the provided target interface.

				Args:
		        payload - post payload
		        target - any type of JSON object in which the JSON response gets stored, passed as &target (pointer to a variable in which response is to be stored) when method is invoked

				Returns nothing if successful, returns error if it fails
	*/
	util.Info(fmt.Sprintf("Posting payload to private endpoint: %s", string(payload)))
	// Ensure required API keys are set
	if os.Getenv("GEMINI_EXCHANGE_API_KEY") == "" || len([]byte(os.Getenv("GEMINI_EXCHANGE_API_SECRET"))) == 0 {
		log.Println("GEMINI_EXCHANGE_API_KEY and GEMINI_EXCHANGE_API_SECRET are not both present in environment")
	} else {
		log.Println("Private credentials from .env successfully loaded for private API functions!")
	}

	// Retrieve API key and secret from environment
	apiKey := util.GetEnvOrDefault("GEMINI_EXCHANGE_API_KEY", "")
	apiSecret := []byte(util.GetEnvOrDefault("GEMINI_EXCHANGE_API_SECRET", ""))

	var payloadJSON map[string]interface{}
	json.Unmarshal(payload, &payloadJSON)
	url := util.GetBaseAPIUrl() + payloadJSON["request"].(string)

	// Base64 encode the JSON payload
	b64Payload := base64.StdEncoding.EncodeToString(payload)

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
	payload, _ := json.Marshal(GetClosedOrdersHistoryRequest{
		Request: "/v1/orders/history",
		Nonce:   util.GenerateNonceString(),
	})
	err := PostPrivateEndpoint(payload, &ordersHistory)
	if err != nil {
		log.Fatalf("Error fetching orders history: %v", err)
		return nil
	}
	return ordersHistory
}

func GetOrderStatus(order_id int) *Order {
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
	payload, _ := json.Marshal(GetOrderStatusRequest{
		OrderID: order_id,
		Request: "/v1/order/status",
		Nonce:   util.GenerateNonceString(),
	})
	err := PostPrivateEndpoint(payload, &orderStatus)

	if err != nil {
		log.Fatalf("Error fetching order status: %v", err)
		return nil
	}
	return &orderStatus
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

	var newOrder Order

	payload, _ := json.Marshal(StopLimitOrderRequest{
		Amount:    strconv.FormatFloat(amount, 'f', 8, 64),
		Price:     strconv.FormatFloat(limitPrice, 'f', 2, 64),
		Side:      "buy",
		StopPrice: strconv.FormatFloat(stopPrice, 'f', 2, 64),
		Symbol:    symbol,
		Type:      "exchange stop limit",
		Request:   "/v1/order/new",
		Nonce:     util.GenerateNonceString(),
	})

	// Pass the payload to the function
	err := PostPrivateEndpoint(payload, &newOrder)

	if err != nil {
		log.Fatalf("Error creating new order: %v", err)
		return nil
	}
	return &newOrder
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

	util.Info(fmt.Sprintf("StopLimitSell called with symbol: %s, amount: %f, stopPrice: %f, limitPrice: %f", symbol, amount, stopPrice, limitPrice))

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

	var newOrder Order

	payload, _ := json.Marshal(StopLimitOrderRequest{
		Amount:    strconv.FormatFloat(amount, 'f', 8, 64),
		Price:     strconv.FormatFloat(limitPrice, 'f', 2, 64),
		Side:      "sell",
		StopPrice: strconv.FormatFloat(stopPrice, 'f', 2, 64),
		Symbol:    symbol,
		Type:      "exchange stop limit",
		Request:   "/v1/order/new",
		Nonce:     util.GenerateNonceString(),
	})

	// Pass the payload to the function
	err := PostPrivateEndpoint(payload, &newOrder)

	if err != nil {
		log.Fatalf("Error creating new order: %v", err)
		return nil
	}
	return &newOrder

}

func GetAvailableBalances() []AvailableBalance {
	util.Info("GetAvailableBalances called")
	var availableBalances []AvailableBalance
	payload, _ := json.Marshal(GetAvailableBalancesRequest{
		Request: "/v1/balances",
		Nonce:   util.GenerateNonceString(),
	})
	err := PostPrivateEndpoint(payload, &availableBalances)
	if err != nil {
		log.Fatalf("Error getting open positions: %v", err)
		return nil
	}
	return availableBalances
}

func filterBalances(balances []AvailableBalance, predicate func(AvailableBalance) bool) []AvailableBalance {
	var result []AvailableBalance
	for _, ab := range balances {
		if predicate(ab) {
			result = append(result, ab)
		}
	}
	return result
}

func GetAvailableCurrencyBalance(currency string) *AvailableBalance {
	util.Info(fmt.Sprintf("GetAvailableBalances, currency: %s", currency))
	availableBalances := GetAvailableBalances()
	util.Info(fmt.Sprintf("My available balances: %v", availableBalances))
	predicate := func(balance AvailableBalance) bool {
		return balance.Currency == currency
	}
	filteredBalances := filterBalances(availableBalances, predicate)
	// filtering by symbol will always return only one position (one position per symbol in your portfolio) so return first item
	if len(filteredBalances) == 0 {
		return nil
	}
	position := filteredBalances[0]
	util.Info(fmt.Sprintf("%s balance: %v", currency, position))
	return &position
}

func CancelOrder(order_id int) *Order {
	util.Info(fmt.Sprintf("CancelOrder called with order_id: %v", order_id))
	var canceledOrder Order
	payload, _ := json.Marshal(CancelOrderRequest{
		Request: "/v1/order/cancel",
		Nonce:   util.GenerateNonceString(),
		OrderID: order_id,
	})
	// Pass the payload to the function
	err := PostPrivateEndpoint(payload, &canceledOrder)
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
