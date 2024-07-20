package main

import "fmt"

func main() {
	var name1 = "Edo"
	var name2 = "Edo"

	var result bool = name1 == name2
	var result2 bool = name1 != name2
	var result3 bool = name1 < name2

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
}