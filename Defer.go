package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}

//Fitur untuk memanggil function ketika eksekusi telah selesai
