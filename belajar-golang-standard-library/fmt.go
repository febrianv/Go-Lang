package main

import "fmt"

func main() {
	fistName, lastName := "Edo", "Febrian"

	//ada space 
	fmt.Println("Hello '", fistName, lastName, "'")
	fmt.Printf("Hello '%s %s'\n", fistName, lastName)
}