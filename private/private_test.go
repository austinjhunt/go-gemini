package private

import (
	"log"
	"testing"
) 


func TestGetOrdersHistory(t *testing.T){
	response := GetOrdersHistory()
	log.Println(response)
	if len(response) == 0 {
		t.Errorf("GetOrdersHistory failed")
	} 
}