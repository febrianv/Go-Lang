package main

import "fmt"

func main() {
	name := "Edo"

	if name == "Edo" {
		fmt.Println("Hallo Edo")
	} else if (name == "Budi") {
		fmt.Println("Hallo Budi")
	} else if (name == "Edo") {
		fmt.Println("Hallo Edo Lagi")
	} else {
		fmt.Println("Hai, boleh kenalan?")
	}



	if 	length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else if (length == 3) {
		fmt.Println("Nama cukup")
	} else {
		fmt.Println("Nama sudah benar")
	}
}