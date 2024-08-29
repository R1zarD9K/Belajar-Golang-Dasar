package main

import "fmt"

func main() {
	const (
		firstName string = "Rahmat"
		lastName = "Pratami"
	)

	// lastName = "Rahmat Hidayat"
	// tipe data constant tidak bisa diubah
	fmt.Println(firstName)
	fmt.Println(lastName)
}