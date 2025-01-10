package main

import (
	"fmt"

	"github.com/austinjhunt/go-gemini/private"
	"github.com/austinjhunt/go-gemini/public"
)
 

func main() {
	// Start the server
	symbols := public.GetSymbols()

	// Print the response
	fmt.Println(symbols)

	private.Test()
}