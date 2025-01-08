package main  

import (  
	"fmt"
	"gogemini/public" 
	"gogemini/private" 
)
 

func main() {
	// Start the server
	symbols := public.GetSymbols()

	// Print the response
	fmt.Println(string(symbols))

	private.Test()
}