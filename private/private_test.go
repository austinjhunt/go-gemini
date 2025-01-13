package private

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/austinjhunt/go-gemini/public"
	"github.com/austinjhunt/go-gemini/util"
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
	lastOrderId, _ := strconv.Atoi(lastOrder.OrderID)
	log.Printf("\nGetting order status for last order in history (id = %v)\n", lastOrderId)
	response := GetOrderStatus(lastOrderId)
	log.Println(response)
	if response == nil {
		t.Errorf("GetOrderStatus failed")
	}
}

func TestStopLimitSell(t *testing.T) {
	log.Println("Testing StopLimitSell function. This will always fail with \"Invalid price for symbol\" if you do not actually have any of that coin to sell.")
	// goal: sell N% of my BTCUSD when BTCUSD ask price reaches some very unlikely low number - sell order will still place as long as i provide an amount of coin to sell that i actually own
	coin := "BTC"
	tradingPair := "btcusd"

	availableBalance := GetAvailableCurrencyBalance(coin)
	// how much I own:
	availableToSell, _ := strconv.ParseFloat(availableBalance.Available, 32)
	if availableToSell == 0 {
		log.Printf("No %s available to sell (might already have an active stop limit sell order)", coin)
	}
	sellRatio := .01 // 1%
	amountToSell := sellRatio * availableToSell

	// don't hardcode the stop and limit prices. get the current price of the coin. use a delta for the purchase prices.
	currentCoinAskPrice, _ := strconv.ParseFloat(public.GetTickerV2(tradingPair).Ask, 32)
	stopPrice := currentCoinAskPrice * .20  // trigger when coin drops to 20% of current value
	limitPrice := currentCoinAskPrice * .15 // do  not accept sell if drops to 15% of current value or below

	order := StopLimitSell(tradingPair, amountToSell, stopPrice, limitPrice)
	if order == nil {
		t.Errorf("StopLimitSell failed: order is nil")
	} else {
		log.Printf("Order successfully placed: %+v", order)
	}

	// cancel that order
	sellOrderId, _ := strconv.Atoi(order.OrderID)
	util.Info("Canceling sell order by ID " + order.OrderID)
	canceledOrder := CancelOrder(sellOrderId)
	util.Info(fmt.Sprintf("Canceled order: %v", canceledOrder))
}
