package private

import (
	"log"
	"strconv"
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
	lastOrderId, _ := strconv.Atoi(lastOrder.OrderID)
	log.Printf("\nGetting order status for last order in history (id = %v)\n", lastOrderId)
	response := GetOrderStatus(lastOrderId)
	log.Println(response)
	if response == nil {
		t.Errorf("GetOrderStatus failed")
	}
}

func TestStopLimitSell(t *testing.T) {
	// Do not run these to avoid flooding account with trade activity and fees.

	// log.Println("Testing StopLimitSell function. This will always fail with \"Invalid price for symbol\" if you do not actually have any of that coin to sell.")
	// // point of a stop limit sell order is to minimize loss, ultimately. sell if it starts dropping and limit your losses with a limit price ("sell if dropping but don't sell if it's already dropped too far")
	// // goal: sell N% of my BTCUSD when BTCUSD ask price reaches some very unlikely low number - sell order will still place as long as i provide an amount of coin to sell that i actually own
	// coin := "BTC"
	// tradingPair := "btcusd"

	// availableBalance := GetAvailableCurrencyBalance(coin)
	// // how much I own:
	// availableToSell, _ := strconv.ParseFloat(availableBalance.Available, 32)
	// if availableToSell == 0 {
	// 	log.Printf("No %s available to sell (might already have an active stop limit sell order)", coin)
	// }
	// sellRatio := .01 // 1%
	// amountToSell := sellRatio * availableToSell

	// // don't hardcode the stop and limit prices. get the current price of the coin. use a delta for the purchase prices.
	// currentCoinAskPrice, _ := strconv.ParseFloat(public.GetTickerV2(tradingPair).Ask, 32)
	// stopPrice := currentCoinAskPrice * .20  // trigger when coin drops to 20% of current value
	// limitPrice := currentCoinAskPrice * .15 // do  not accept sell if drops to 15% of current value or below

	// order := StopLimitSell(tradingPair, amountToSell, stopPrice, limitPrice)
	// if order == nil {
	// 	t.Errorf("StopLimitSell failed: order is nil")
	// } else {
	// 	log.Printf("Order successfully placed: %+v", order)
	// }

	// // cancel that order
	// sellOrderId, _ := strconv.Atoi(order.OrderID)
	// util.Info("Canceling sell order by ID " + order.OrderID)
	// canceledOrder := CancelOrder(sellOrderId)
	// util.Info(fmt.Sprintf("Canceled order: %v", canceledOrder))
}

func TestStopLimitBuy(t *testing.T) {

	// Do not run these to avoid flooding account with trade activity and fees.

	// log.Println("Testing StopLimitBuy function.")
	// // point of a stop limit buy order is to trigger a buy when you see a spike but to limit your expense with a limit price, i.e. buy if it reaches $100K but if it's already $103K (limit) then do not buy
	// tradingPair := "btcusd"

	// // how much are we buying? $50 worth..
	// spendUSDAmount := 50.00
	// amountToBuy := public.ConvertUSDToCryptoAmount(spendUSDAmount, tradingPair)

	// // don't hardcode the stop and limit prices. get the current price of the coin. use a delta for the purchase prices.
	// currentCoinAskPrice, _ := strconv.ParseFloat(public.GetTickerV2(tradingPair).Ask, 32)
	// stopPrice := currentCoinAskPrice * 1.40 // trigger when coin spikes to 140% of current value
	// limitPrice := currentCoinAskPrice * 1.5 // do not buy if 150% or higher

	// order := StopLimitBuy(tradingPair, amountToBuy, stopPrice, limitPrice)
	// if order == nil {
	// 	t.Errorf("StopLimitSell failed: order is nil")
	// } else {
	// 	log.Printf("Order successfully placed: %+v", order)
	// }

	// // cancel that order
	// buyOrderId, _ := strconv.Atoi(order.OrderID)
	// util.Info("Canceling buy order by ID " + order.OrderID)
	// canceledOrder := CancelOrder(buyOrderId)
	// util.Info(fmt.Sprintf("Canceled order: %v", canceledOrder))
}
