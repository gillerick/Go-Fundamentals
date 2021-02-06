package main

import "fmt"

type Stocks struct {
	Symbol string
	Price  float64
	Volume int
	Buy    bool
}

func (t *Stocks) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value
}

func main() {
	t1 := Stocks{"MSFT", 242.02, 89, false}
	fmt.Println(t1.Value())

}
