package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Rahmat"
	names[1] = "Pratami"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [...]int{
		90,
		80,
	}
	fmt.Println(values)
	fmt.Println(len(values))
}
