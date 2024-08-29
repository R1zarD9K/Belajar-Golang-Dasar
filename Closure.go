package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Incrementing counter")
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println("Counter:", counter)
}

//Hati-hati dalam menggunakan increment jika kode sudah sangat kompleks
