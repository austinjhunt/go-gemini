package public 

import (
	"fmt" 
	"time"
)

type TickerV1 struct {  
	Ask 	string 	`json:"ask"`
	Bid 	string	`json:"bid"`
	Last	string	`json:"last"`
	Volume 	map[string]interface{} 
}  

type TickerV2 struct {
	Symbol	string		`json:"symbol"`
	Open 	string 		`json:"open"`
	High 	string 		`json:"high"`
	Low 	string 		`json:"low"`
	Close 	string		`json:"close"`
	Changes []string  	`json:"changes"`
	Bid 	string		`json:"bid"`
	Ask 	string 		`json:"ask"`
}
 

func PrintTickerV1(ticker TickerV1) {
	fmt.Println("TickerV1: ")
	fmt.Printf("  Last: %s\n", ticker.Last)
	fmt.Printf("  Bid: %s\n", ticker.Bid)
	fmt.Printf("  Ask: %s\n", ticker.Ask)
	fmt.Printf("  Volume: \n")
	for k, v := range ticker.Volume { 
		if (k == "timestamp") {
			timestamp := time.UnixMilli(int64(v.(float64)))
			fmt.Printf("    %s: %v\n", k, timestamp.Format(time.RFC1123))
		} else {
			fmt.Printf("    %s: %v\n", k, v)
		} 
	}
}

func PrintTickerV2(ticker TickerV2) {
	fmt.Println("TickerV2: ")
	fmt.Printf("  Symbol: %s\n", ticker.Symbol)
	fmt.Printf("  Open: %s\n", ticker.Open)
	fmt.Printf("  High: %s\n", ticker.High)
	fmt.Printf("  Low: %s\n", ticker.Low)
	fmt.Printf("  Close: %s\n", ticker.Close) 
	fmt.Printf("  Bid: %s\n", ticker.Bid)
	fmt.Printf("  Ask: %s\n", ticker.Ask)
	fmt.Printf("  Changes (hourly, past 24h): \n")
	for k, v := range ticker.Changes { 
		fmt.Printf("    %dh ago: %v\n", k+1, v)
	}
}
