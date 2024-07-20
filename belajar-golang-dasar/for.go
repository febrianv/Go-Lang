package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	fmt.Println("Selesai")

	// manual
	names := []string{"Edo", "Febrian", "Vernando"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	// for range with out index/key
	for _, name := range names {
		fmt.Println(name)
	}
}