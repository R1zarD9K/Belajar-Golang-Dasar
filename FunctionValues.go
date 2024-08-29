package main

import "fmt"

func Goodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	kontol := Goodbye
	anjing := Goodbye
	fmt.Println(kontol("Rahmat"))
	fmt.Println(anjing("Pratami"))
}