package main

import ( 
	"fmt"
	"github.com/austinjhunt/go-gemini/public"
	"github.com/austinjhunt/go-gemini/util"
	"github.com/manifoldco/promptui" // allow interactive scrolling & selection of options 
)

func main() {
	introBanner:= `
	$$\      $$\ $$$$$$$$\ $$\        $$$$$$\   $$$$$$\  $$\      $$\ $$$$$$$$\          
	$$ | $\  $$ |$$  _____|$$ |      $$  __$$\ $$  __$$\ $$$\    $$$ |$$  _____|         
	$$ |$$$\ $$ |$$ |      $$ |      $$ /  \__|$$ /  $$ |$$$$\  $$$$ |$$ |               
	$$ $$ $$\$$ |$$$$$\    $$ |      $$ |      $$ |  $$ |$$\$$\$$ $$ |$$$$$\             
	$$$$  _$$$$ |$$  __|   $$ |      $$ |      $$ |  $$ |$$ \$$$  $$ |$$  __|            
	$$$  / \$$$ |$$ |      $$ |      $$ |  $$\ $$ |  $$ |$$ |\$  /$$ |$$ |               
	$$  /   \$$ |$$$$$$$$\ $$$$$$$$\ \$$$$$$  | $$$$$$  |$$ | \_/ $$ |$$$$$$$$\          
	\__/     \__|\________|\________| \______/  \______/ \__|     \__|\________|         
																						
																						
																						
	$$$$$$$$\  $$$$$$\                                                                   
	\__$$  __|$$  __$$\                                                                  
	$$ |   $$ /  $$ |                                                                 
	$$ |   $$ |  $$ |                                                                 
	$$ |   $$ |  $$ |                                                                 
	$$ |   $$ |  $$ |                                                                 
	$$ |    $$$$$$  |                                                                 
	\__|    \______/                                                                  
																						
																						
																						
	$$$$$$\   $$$$$$\         $$$$$$\  $$$$$$$$\ $$\      $$\ $$$$$$\ $$\   $$\ $$$$$$\ 
	$$  __$$\ $$  __$$\       $$  __$$\ $$  _____|$$$\    $$$ |\_$$  _|$$$\  $$ |\_$$  _|
	$$ /  \__|$$ /  $$ |      $$ /  \__|$$ |      $$$$\  $$$$ |  $$ |  $$$$\ $$ |  $$ |  
	$$ |$$$$\ $$ |  $$ |      $$ |$$$$\ $$$$$\    $$\$$\$$ $$ |  $$ |  $$ $$\$$ |  $$ |  
	$$ |\_$$ |$$ |  $$ |      $$ |\_$$ |$$  __|   $$ \$$$  $$ |  $$ |  $$ \$$$$ |  $$ |  
	$$ |  $$ |$$ |  $$ |      $$ |  $$ |$$ |      $$ |\$  /$$ |  $$ |  $$ |\$$$ |  $$ |  
	\$$$$$$  | $$$$$$  |      \$$$$$$  |$$$$$$$$\ $$ | \_/ $$ |$$$$$$\ $$ | \$$ |$$$$$$\ 
	\______/  \______/        \______/ \________|\__|     \__|\______|\__|  \__|\______|
																																																																											
	`
	fmt.Println(introBanner)

	// Start the server
	symbols := public.GetSymbols()

	prompt := promptui.Select{
        Label: "Choose a ticker symbol to get details! (left and right to jump between pages)",
        Items: symbols,
        Size: 10, // number of items shown at a time
		
    }


	selectedSymbolIndex, selectedSymbol, err := prompt.Run()

    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        return
    }

    fmt.Printf("You chose #%d: %s\n", selectedSymbolIndex+1, selectedSymbol)

	 

	prompt = promptui.Select{
		Label: "What would you like to do with the symbol " + selectedSymbol,
		Items: []string{"Get symbol details", "Get ticker", "Get ticker v2"},
		Size: 3,  
	}


	actionIndex, actionResult, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You chose #%d: %s\n", actionIndex+1, actionResult)
	 
	if (actionResult == "Get symbol details") {
		fmt.Println("Getting symbol details...")
		util.PrintMap(public.GetSymbolDetails(selectedSymbol))
	} else if (actionResult == "Get ticker") {
		fmt.Println("Getting ticker...")
		public.PrintTickerV1(*public.GetTicker(selectedSymbol))
	} else if (actionResult == "Get ticker v2") { 
		fmt.Println("Getting ticker v2...")
		public.PrintTickerV2(*public.GetTickerV2(selectedSymbol))
	}  
}

 