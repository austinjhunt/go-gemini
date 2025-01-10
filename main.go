package main

import (
	"log"

	"github.com/austinjhunt/go-gemini/public"
)

func main() {

	// Start the server
	symbols := public.GetSymbols()

	// Print the response
	log.Println(symbols) 
}

 