package public

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/austinjhunt/go-gemini/util"
	"github.com/xuri/excelize/v2"
)

// TestGetSymbols tests the GetSymbols function
func TestGetSymbols(t *testing.T) {
	response := GetSymbols()
	log.Println(response)
	// response should be an array of supported symbols
	if len(response) == 0 {
		t.Errorf("GetSymbols failed")
	}
	// example symbols that should be in the response: btcusd, ethusd, ethbtc, bchusd, bchbtc, bcheth, ltcusd, ltcbtc, ltceth, ltcbch, batusd, daiusd, linkusd, oxtusd, linkbtc, linketh, ampusd, compusd, paxgus
	expectedSymbols := []string{"btcusd", "ethusd", "ethbtc", "bchusd", "bchbtc", "bcheth", "ltcusd", "ltcbtc", "ltceth", "ltcbch", "batusd", "daiusd", "linkusd", "oxtusd", "linkbtc", "linketh", "ampusd", "compusd", "paxgusd"}
	// if response does not contain all expected symbols, test failed
	for _, symbol := range expectedSymbols {
		// if response does not contain the symbol, test failed
		if !util.ArrayContainsString(response, symbol) { 
			t.Errorf("GetSymbols failed")
		}
	} 
}

// TestGetSymbolDetails tests the GetSymbolDetails function
func TestGetSymbolDetails(t *testing.T) {
	response := GetSymbolDetails("btcusd")
	log.Println(response)
	if len(response) == 0 {
		t.Errorf("GetSymbolDetails failed")
	}
}

// TestGetNetwork tests the GetNetwork function
func TestGetNetwork(t *testing.T) {
	response := GetNetwork("btc")
	log.Println(response)
	if len(response) == 0 {
		t.Errorf("GetNetwork failed")
	}
}

// TestGetTicker tests the GetTicker function
func TestGetTicker(t *testing.T) {
	response := GetTicker("btcusd")
	log.Println(response)
	if len(response) == 0 {
		t.Errorf("GetTicker failed")
	}
}

// TestGetTickerV2 tests the GetTickerV2 function
func TestGetTickerV2(t *testing.T) {
	response := GetTickerV2("btcusd")
	log.Println(response)
	if len(response) == 0 {
		t.Errorf("GetTickerV2 failed")
	}
}

// TestGetCandles tests the GetCandles function
func TestGetCandles(t *testing.T) {
	response := GetCandles("btcusd", "1m")
	log.Println(response)
	if len(response) == 0 {
		t.Errorf("GetCandles failed")
	}
}
 

func TestGetDerivativesCandles(t *testing.T) {
  response := GetDerivativesCandles("btcusd", "1m")
  log.Println(response) 
  if len(response) == 0 {
    t.Errorf("GetDerivativesCandles failed")
  }
}

func TestGetFeePromos(t *testing.T) {
  response := GetFeePromos()
  log.Println(response)
  if len(response) == 0 {
    t.Errorf("GetFeePromos failed")
  }
}

func TestGetCurrentOrderBook(t *testing.T) {
  response := GetCurrentOrderBook("btcusd") 
  log.Println(response) 
  if len(response) == 0 {
    t.Errorf("GetCurrentOrderBook failed")
  }
}

func TestGetTradeHistor(t *testing.T){
  response := GetTradeHistory("btcusd")
  log.Println(response)
  if len(response) == 0 {
    t.Errorf("GetTradeHistory failed")
  }
}

func TestGetPriceFeed(t *testing.T){
  response := GetPriceFeed()
  log.Println(response)
  if len(response) == 0 {
    t.Errorf("GetPriceFeed failed")
  }
}



func TestGetFundingAmount(t *testing.T){
  response := GetFundingAmount("BTCGUSDPERP")
  log.Println(response)
  if len(response) == 0 {
    t.Errorf("GetFundingAmount failed")
  }
}


func TestDownloadFundingAmountReport(t *testing.T) {
	// Create a mock XLSX file in memory with the required columns.
	mockFile := excelize.NewFile()
	sheetName := mockFile.GetSheetName(0)

	// Write column headers.
	headers := []string{"Symbol", "FundingDateTime", "Amount", "Is Realized ?"}
	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1)
		mockFile.SetCellValue(sheetName, cell, header)
	}

	// Save the mock file to a buffer.
	var buf bytes.Buffer
	if err := mockFile.Write(&buf); err != nil {
		t.Fatalf("failed to create mock XLSX file: %v", err)
	}

	// Start a mock HTTP server to serve the mock XLSX file.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(buf.Bytes())
	}))
	defer server.Close()

  
	mockSymbol := "BTCGUSDPERP"

	// Calculate mockFromDate and mockToDate for the past week.
	now := time.Now()
	mockToDate := now.Format("2006-01-02")
	mockFromDate := now.AddDate(0, 0, -7).Format("2006-01-02")
	mockNumRows := 100

	// Call the function to be tested.
	err := DownloadFundingAmountReport(mockSymbol, mockFromDate, mockToDate, mockNumRows)
	if err != nil {
		t.Fatalf("DownloadFundingAmountReport failed: %v", err)
	}

	// Validate the downloaded file.
	expectedFileName := "funding_amount_report_" + mockSymbol + "_" + mockFromDate + "_to_" + mockToDate + ".xlsx"
	defer os.Remove(expectedFileName) // Clean up after the test.

	// Check if the file exists.
	if _, err := os.Stat(expectedFileName); os.IsNotExist(err) {
		t.Fatalf("file was not downloaded: %s", expectedFileName)
	}

	// Open the downloaded file.
	downloadedFile, err := excelize.OpenFile(expectedFileName)
	if err != nil {
		t.Fatalf("failed to open downloaded file: %v", err)
	}

	// Get the first sheet name from the downloaded file.
	sheetNames := downloadedFile.GetSheetList()
	if len(sheetNames) == 0 {
		t.Fatalf("no sheets found in the downloaded file")
	}
	firstSheetName := sheetNames[0] // Use the first sheet name.

	// Validate column headers in the first row.
	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1)
		value, _ := downloadedFile.GetCellValue(firstSheetName, cell)
		if value != header {
			t.Errorf("expected header %q, got %q", header, value)
		}
	}

	t.Log("TestDownloadFundingAmountReport passed.")
}