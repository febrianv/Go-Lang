package main

import "fmt"

type Address struct {
	City, Province, Country string
}


func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // pinter

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{"Jakart", "DKI Jakart", "Indonesia"}

	fmt.Println(address1) // address 1 ikut berubah sama seperti address 2
	fmt.Println(address2)
}