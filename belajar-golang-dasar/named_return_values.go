package main

import "fmt"

// func getCompleteName() (firtName string, middleName string, lastName string) {
func getCompleteName() (firtName, middleName, lastName string) {
	firtName = "Edo"
	// middleName = "Febrian"
	// lastName = "Vernando"

	return firtName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()

	fmt.Println(a, b, c)
}