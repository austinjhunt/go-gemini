package private

import (
	"log"

	"github.com/austinjhunt/go-gemini/util"
)
 
 
func GetOrdersHistory() []map[string]interface{} {
	/* 
	This API retrieves (closed) orders history for an account.

	The API key you use to access this endpoint must have the Trader or Auditor role assigned. See Roles for more information.
	*/ 

	var ordersHistory []map[string]interface{}
	// Create the payload map separately
	payload := map[string]interface{}{}

	endpoint := "/v1/orders/history"

	// Pass the payload to the function
	err := util.PostPrivateEndpoint(endpoint, payload, &ordersHistory)

	if err != nil {
		log.Fatalf("Error fetching orders history: %v", err)
		return nil 
	}
	return ordersHistory
}
 
func GetOrderStatus(order_id int, include_trades bool, account string) map[string]interface{} {
	/* 
	Get order status

	Args: 
	order_id (int) - The order id to get information on. The order_id represents a whole number and is transmitted as an unsigned 64-bit integer in JSON format. An unsigned 64-bit integer is a non-negative whole number with a maximum value of 18,446,744,073,709,551,615. order_id cannot be used in combination with client_order_id.
	include_trades (bool) - if True the endpoint will return individual trade details of all fills from the order. 
	account (string) - Required for Master API keys as described in Private API Invocation. The name of the account within the subaccount group. Specifies the account on which the order was placed. Only available for exchange accounts. Pass empty string if not using master API key. 
	*/ 
	
	var orderStatus map[string]interface{}
	endpoint := "/v1/order/status"

	// Create the payload map separately
	payload := map[string]interface{}{
		"order_id":      order_id,
		"include_trades": include_trades,
		"account":       account,
	}

	// Pass the payload to the function
	err := util.PostPrivateEndpoint(endpoint, payload, &orderStatus)

	if err != nil {
		log.Fatalf("Error fetching order status: %v", err)
		return nil 
	}
	return orderStatus



}