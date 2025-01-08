package main  

import (  
	"fmt"
	"github.com/austinjhunt/go-gemini/public" 
	"github.com/austinjhunt/go-gemini/private" 
)
 

func main() {
	// Start the server
	symbols := public.GetSymbols()

	// Print the response
	fmt.Println(string(symbols))

	private.Test()
}