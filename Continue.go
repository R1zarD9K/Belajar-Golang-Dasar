package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 { // Jika i genap, maka continue
			continue
		}
		fmt.Println("Angka ganjil ke-", i)
	}
}
