package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blackList Blacklist) {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Edo", blackList)

	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}