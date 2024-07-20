package main

import "fmt"

func getFullName() (string, string) {
	return "Febrian", "Vernando"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
}