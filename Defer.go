package main

import "fmt"

//Defer is used to close resources such as files, virtual machines, sockets

func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func worker() {
	defer cleanup("A")
	fmt.Println("Worker")

}

func main() {
	worker()
}
