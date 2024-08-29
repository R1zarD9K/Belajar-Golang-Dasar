package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Rahmat Pratami",
		"address": "Indonesia",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := map[string]string{
		"title":  "Belajar Golang",
		"author": "Rahmat Pratami",
		"ups":    "Salah",
	}
	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
	
}
