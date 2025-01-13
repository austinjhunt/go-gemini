package private

type Order struct {
	OrderID           string   `json:"order_id"`
	ID                string   `json:"id"`
	Symbol            string   `json:"symbol"`
	Exchange          string   `json:"exchange"`
	AvgExecutionPrice string   `json:"avg_execution_price"`
	Side              string   `json:"side"`
	Type              string   `json:"type"`
	Timestamp         string   `json:"timestamp"`
	TimestampMs       int64    `json:"timestampms"`
	IsLive            bool     `json:"is_live"`
	IsCancelled       bool     `json:"is_cancelled"`
	IsHidden          bool     `json:"is_hidden"`
	WasForced         bool     `json:"was_forced"`
	ExecutedAmount    string   `json:"executed_amount"`
	Options           []string `json:"options"`
	StopPrice         string   `json:"stop_price"`
	Price             string   `json:"price"`
	OriginalAmount    string   `json:"original_amount"`
}

type GetClosedOrdersHistoryRequest struct {
	Request string `json:"request"`
	Nonce   string `json:"nonce"`
}

type GetOrderStatusRequest struct {
	OrderID int    `json:"order_id"`
	Request string `json:"request"`
	Nonce   string `json:"nonce"`
}

type StopLimitOrderRequest struct {
	Amount    string `json:"amount"`
	Price     string `json:"price"`
	Side      string `json:"side"`
	StopPrice string `json:"stop_price"`
	Symbol    string `json:"symbol"`
	Type      string `json:"type"`
	Request   string `json:"request"`
	Nonce     string `json:"nonce"`
}

type CancelOrderRequest struct {
	OrderID int    `json:"order_id"`
	Request string `json:"request"`
	Nonce   string `json:"nonce"`
}

type AvailableBalance struct {
	Type                   string `json:"type"`
	Currency               string `json:"currency"`
	Amount                 string `json:"amount"`
	Available              string `json:"available"`
	AvailableForWithdrawal string `json:"availableForWithdrawal"`
}

type GetAvailableBalancesRequest struct {
	Request string `json:"request"`
	Nonce   string `json:"nonce"`
}
