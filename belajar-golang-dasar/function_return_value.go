package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Edo")
	fmt.Println(result)
	fmt.Println(getHello("Budi"))
}