package main

import "fmt"

func main() {
	//Declare book variable
	book := "The colour of magic"

	//Return book
	fmt.Println(book)
	//Return length of book
	fmt.Println(len(book))
	//Return book object and type
	fmt.Printf("Book = %v (type = %T)", book[0], book[0])
	fmt.Println("t" + book[1:])

	//Strings ig go are immutable

}
