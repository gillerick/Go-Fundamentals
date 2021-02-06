package main

import "fmt"

//Fields that start with uppercase are accessible from other packages, everything else is private to a package
type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func main() {
	t1 := Trade{"MSFT", 180, 242.01, true}

	//Alternative method of creating struct objects - specifying field names
	//If ones uses this method, they don't have to follow the struct fields order
	t2 := Trade{Symbol: "AMZN", Price: 3090.90, Buy: false, Volume: 12}
	fmt.Printf("%+v\n", t2)
	//The preceding + symbol gives us a detailed description of the struct object
	fmt.Printf("%+v\n", t1)
	fmt.Printf("%v", t1)
	//fmt.Println(t1)

}
