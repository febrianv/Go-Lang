package main

import "fmt"

type Man struct {
	Name string
}

//// ini tidak mengubah data aslinya ketika dipanggil
// func (man Man) Merried() {
// 	man.Name = "Mr. " + man.Name
// }

// yang dikirim hanya pointernya 
// jadi ketika dirubah didalam method maka data aslinya ikut berubah
func (man *Man) Merried() {
	man.Name = "Mr. " + man.Name
}

func main() {
	eko := Man{"Eko"}
	eko.Merried()

	fmt.Println(eko)
}