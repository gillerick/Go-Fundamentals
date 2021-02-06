package main

import (
	"fmt"
	"os"
)

type StockTrade struct {
	Symbol string
	Price  float64
	Volume int
	Buy    bool
}

func NewStockTrade(symbol string, price float64, volume int, buy bool) (*StockTrade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Symbol can't be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("Volume must be greater than 0 (was %d)", volume)
	}

	if price <= 0.0 {
		return nil, fmt.Errorf("Price must be greater than 0 (was %f)", price)
	}

	stockTrade := &StockTrade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}

	return stockTrade, nil

}

func (t *StockTrade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func main() {
	t, err := NewStockTrade("MSFT", 243.2, 34, true)
	if err != nil {
		fmt.Printf("Error: %s, can't create trade", err)
		os.Exit(1)
	}
	fmt.Println(t.Value())
}
