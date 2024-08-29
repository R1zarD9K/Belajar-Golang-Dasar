package main

import "fmt"

func main(){
	type NoKTP string

	var ktpRahmat NoKTP = "1234567890"

	var contoh string = "2222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpRahmat)
	// fmt.Println(contoh)
	fmt.Println(contohKtp)
}
//type fungsinya adalah untuk membuat alias dari suatu tipe data