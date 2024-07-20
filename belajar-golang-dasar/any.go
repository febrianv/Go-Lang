package main

import "fmt"

func Ups() any {
	// return 1
	// return true
	return "Berhasil"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}