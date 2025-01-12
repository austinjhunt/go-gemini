package private

type Order struct {
	OrderID            string   `json:"order_id"`
	ID                 string   `json:"id"`
	Symbol             string   `json:"symbol"`
	Exchange           string   `json:"exchange"`
	AvgExecutionPrice  string   `json:"avg_execution_price"`
	Side               string   `json:"side"`
	Type               string   `json:"type"`
	Timestamp          string    `json:"timestamp"`
	TimestampMs        int64    `json:"timestampms"`
	IsLive             bool     `json:"is_live"`
	IsCancelled        bool     `json:"is_cancelled"`
	IsHidden           bool     `json:"is_hidden"`
	WasForced          bool     `json:"was_forced"`
	ExecutedAmount     string   `json:"executed_amount"`
	Options            []string `json:"options"`
	StopPrice          string   `json:"stop_price"`
	Price              string   `json:"price"`
	OriginalAmount     string   `json:"original_amount"`
}

type GetOrderStatusOptions struct {
	ClientOrderID string `json:"client_order_id"`
	Account string `json:"account"`
	IncludeTrades bool `json:"include_trades"`
}


type NewOrderOptions struct {
	Options []interface{}
	ClientOrderID string `json:"client_order_id"`
	StopPrice string `json:"stop_price"`
	Account string `json:"account"`
} 


type StopLimitOrderOptions struct {
	StopPrice string `json:"stop_price"` 
} 