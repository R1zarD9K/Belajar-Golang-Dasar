package main

import "fmt"

func main() {
	var name1 = "Rahmat"
	var name2 = "Rahmat"

	var result = name1 == name2
	fmt.Println(result)
	
	result = name1 != name2
	fmt.Println(result)
}