package main

import "fmt"

func random() any {
	return 100
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// var resultInt = result.(int) // panic
	// fmt.Println(resultInt)

	
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)		
	}
}