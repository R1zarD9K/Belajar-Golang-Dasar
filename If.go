package main

import "fmt"

func main() {
	// name := "Rahmat Pratami"
	// name := "Rahmat"
	name := "Kontol"

	if name == "Rahmat Pratami" {
		fmt.Println("Hello Rahmat Pratami")
	} else if name == "Rahmat" {
		fmt.Println("Hello Rahmat")
	} else {
		fmt.Println("Lo siapa goblok?")
	}

	length := len(name)
	if length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
	 
}
