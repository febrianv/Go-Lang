package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{}
	person["name"] = "Edo"
	person["address"] = "Bandung"

	person2 := map[string]string{
		"name":    "Edo",
		"address": "Bandung",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["salah"]) // akan mereturn default value

	fmt.Println(person2)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Febrian"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)

	delete(book, "asal") //walaupun datanya tidak ada, tidak terjadi error
	fmt.Println(book)
}