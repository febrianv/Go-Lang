package main

import "fmt"

func sayHalloTo(firtName string, lastName string) {
	fmt.Println("Hello", firtName, lastName)
}

func main() {
	// sayHalloTo() //error
	// sayHalloTo(1,2) // error
	sayHalloTo("Febrian", "Vernando")
	sayHalloTo("Budi", "Nugraha")
}