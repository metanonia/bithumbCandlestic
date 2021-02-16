package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println(os.Args[0], "order_currency  interval" )
		os.Exit(1)
	}
	var path = "https://api.bithumb.com/public/candlestick/"+args[0]+"/"+args[1]
	resp, err := http.Get(path)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var dat map[string] interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	status :=  dat["status"].(string)
	if status == "0000" {
		list := dat["data"].([]interface{})
		for _, item := range list {
			str := fmt.Sprintf("%v", item)
			var (
				baseTime 	float64
				startPrice 	uint64
				endPrice	uint64
				highPrice 	uint64
				lowPrice	uint64
				tradingVol	float64
			)

			fmt.Sscanf(str, "[%e %d %d %d %d %f", &baseTime, &startPrice, &endPrice, &highPrice, &lowPrice, &tradingVol)
			uTime := time.Unix(int64(baseTime)/1000,0)
			fmt.Printf("%s, %d, %d, %d, %d, %f\n", uTime.Format("2006/01/01 15:04:05"), startPrice, endPrice, highPrice, lowPrice, tradingVol)
		}
	} else {
		fmt.Printf("Path: %s\n",path)
		fmt.Printf("Error(%s): %s\n",status, dat["message"].(string))
	}
}
