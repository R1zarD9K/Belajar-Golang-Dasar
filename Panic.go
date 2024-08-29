package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover() //Cara penangan error untuk mengatasi langsung mati saat panic
	fmt.Println("Terjadi panic :", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APP ERROR") //Cara untuk menghentikan function ketika error
	}

}
func main() {
	// runApp(false)
	runApp(true)
	fmt.Println("Test panic")
}
