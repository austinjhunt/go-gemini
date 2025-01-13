package public

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