package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"AMZN": 3331.,
		"GOOG": 2062.37,
		"MSFT": 242.01,
		"AAPL": 137.39,
		"TSLA": 849.99, //Must have a trailing comma
	}

	fmt.Printf("Microsoft stocks as of 16:38 Friday 5th Feb were %v\n\n", stocks["MSFT"])

	value, ok := stocks["TLA"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Printf("%v: Stock not found\n\n", ok)
	}

	//Set map value
	stocks["MSFT"] = 243
	fmt.Println(stocks["MSFT"])

	//Delete map item
	delete(stocks, "AMZN")

	//Loop through map items
	for k, v := range stocks {
		fmt.Println(k, v)
	}

}
