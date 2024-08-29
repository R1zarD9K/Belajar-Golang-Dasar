package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Asu"
	}
	registerUser("Rahmat", blacklist)
	registerUser("Asu", blacklist)

	// fmt.Println("==========Pilih cara atas atau bawah sama saja===========")

	// registerUser("Rahmat", func(name string) bool {
	// 	return name == "Asu"
	// })
	// registerUser("Asu", func(name string) bool {
	// 	return name == "Asu"
	// })
}
