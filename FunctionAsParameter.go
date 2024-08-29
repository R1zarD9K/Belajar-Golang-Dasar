package main

import "fmt"

type Filter func(string) string // Filter is a function type for alias

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Kontol" {
		return "..."
	} else {
		return name
	}
}

func main() {
	filter := spamFilter
	sayHelloWithFilter("Rahmat", filter)
	sayHelloWithFilter("Kontol", filter)
}
