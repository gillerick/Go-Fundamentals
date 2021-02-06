package main

import "fmt"

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func main() {
	t1 := Trade{"MSFT", 18000000, 242.01, true}
	//The preceding + symbol gives us a detailed description of the struct object
	fmt.Printf("%+v\n", t1)
	fmt.Printf("%v", t1)
	//fmt.Println(t1)

}
