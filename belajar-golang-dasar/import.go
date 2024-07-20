package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Edo")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(helper.version) // tidak bisa di akses
	fmt.Println(helper.sayGoodBye("edo")) // tidak bisa di akses

}