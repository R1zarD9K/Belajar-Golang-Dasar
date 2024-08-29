package main

import "fmt"

func getFullName() (string, string) {
	return "John", "Doe"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName)
	// fmt.Println(lastName)
	// fmt.Println(firstName, lastName)
	firstName, _ := getFullName()
	fmt.Println(firstName, _)
}
