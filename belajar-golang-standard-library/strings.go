package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Febrian Vernando", "Febri"))
	fmt.Println(strings.Split("Febrian Vernando", " "))
	fmt.Println(strings.ToLower("Febrian Vernando"))
	fmt.Println(strings.ToUpper("Febrian Vernando"))
	fmt.Println(strings.Trim("           Febrian Vernando            ", " "))
	fmt.Println(strings.ReplaceAll("Febrian Vernando", "Febri", "Wawa"))
}