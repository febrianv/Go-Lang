package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 *Address = &Address{}
	ChangeCountryToIndonesia(address1)

	var address2 Address = Address{}
	ChangeCountryToIndonesia(&address2)

	fmt.Println(address1)
	fmt.Println(address2)
}