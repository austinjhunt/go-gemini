package private

import (
	"log"
	"testing"
)

func TestGetClosedOrdersHistory(t *testing.T) {
	log.Println("Getting closed orders history")

	response := GetClosedOrdersHistory()

	log.Println(response)
	if response == nil {
		t.Errorf("GetOrdersHistory failed")
	}
}

func TestGetOrderStatus(t *testing.T) {
	log.Println("Getting order status")

	ordersHistory := GetClosedOrdersHistory()
	lastOrder := ordersHistory[len(ordersHistory)-1]
	lastOrderId := lastOrder.OrderID
	log.Printf("\nGetting order status for last order in history (id = %v)\n", lastOrderId)
	response := GetOrderStatus(lastOrderId, nil)
	log.Println(response)
	if response == nil {
		t.Errorf("GetOrderStatus failed")
	}
}

// should add a test for StopLimitBuy and StopLimitSell
func TestStopLimitBuy(t *testing.T) {
	log.Println("Testing StopLimitBuy function")

	symbol := "BTCUSD"
	dollarAmount := 100.0
	stopPrice := 45000.0
	limitPrice := 45100.0

	order := StopLimitBuy(symbol, dollarAmount, stopPrice, limitPrice)
	if order == nil {
		t.Errorf("StopLimitBuy failed: order is nil")
	} else {
		log.Printf("Order successfully placed: %+v", order)
	}
}

func TestStopLimitSell(t *testing.T) {
	log.Println("Testing StopLimitSell function")

	symbol := "BTCUSD"
	amount := 0.002
	stopPrice := 999000.0
	limitPrice := 98900.0

	order := StopLimitSell(symbol, amount, stopPrice, limitPrice)
	if order == nil {
		t.Errorf("StopLimitSell failed: order is nil")
	} else {
		log.Printf("Order successfully placed: %+v", order)
	}
}
