package main

import (
	"encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
	"time"
)

type CoinbaseResponse struct {
	Price    string    `json:"price"`
}
type BitflyerResponse struct {
	Price    float64    `json:"best_bid"` // using best_bid, no price from api
}

func main(){
	// get price of BTC for 1 minute, update ever 2 seconds
	for i := 0; i < 30; i++{
		getCoinbasePrice("BTC-USD")
		getBitflyerPrice("BTC_USD")
		
		fmt.Println()
		
		time.Sleep(2 * time.Second)
	}
	
}


func getCoinbasePrice(product string){
	
	// create coinbase endpoint for product ticker
	endpoint := "https://api.pro.coinbase.com/products/" + product + "/ticker"
	
	// get json repsonse from endpoint
	response, err := http.Get(endpoint)
	
	// kill on error
	if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    //fmt.Println(string(responseData))
	
	var responseObject CoinbaseResponse
	json.Unmarshal(responseData, &responseObject)
	
	fmt.Println("Current " + product + " price via Coinbase: ", responseObject.Price)
}


func getBitflyerPrice(product string){
	endpoint := "https://api.bitflyer.com/v1/getticker?product_code=" + product
	
	response, err := http.Get(endpoint)
	if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }
	
	responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
	//fmt.Println(string(responseData))
	
	var responseObject BitflyerResponse
	json.Unmarshal(responseData, &responseObject)
	
	fmt.Println("Current " + product + " price via Bitfly: ", floatToString(responseObject.Price))
}


func floatToString(float float64) string{
	s := fmt.Sprintf("%.2f", float)
	return s
}