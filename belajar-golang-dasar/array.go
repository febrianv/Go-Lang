package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Febrian"
	name[1] = "Vernando"
	name[2] = "Edo"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [...]int {
		90,
		88,
		95,
		100,
		110,
	}

	fmt.Println(len(values))

	values[0] = 10
	fmt.Println(values)
	
}