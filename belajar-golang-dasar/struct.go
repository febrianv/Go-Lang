package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
} 

func main() {
	var edo Customer
	fmt.Println(edo) // akan dibalikan default value

	edo.Name = "Febrian Vernando"
	edo.Address = "Bandung"
	edo.Age = 17

	fmt.Println(edo)
	fmt.Println(edo.Name)
	fmt.Println(edo.Address)
	fmt.Println(edo.Age)

	joko := Customer{
		Name: "Joko Anwar",
		Address: "Yogyakart",
		Age: 50,
	}

	fmt.Println(joko)

	budi := Customer {"Budi Utomo", "Padang", 100}
	fmt.Println(budi)

	budi.sayHello("Agus")
}
