package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 //pointer

	address2.City = "Kuala Lumpur"
	address2.Country = "Malaysia"
	fmt.Println(address1) //ikut berubah
	fmt.Println(address2) //berubah

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} //Ini disebut operater asterisk
	fmt.Println(address1)
	fmt.Println(address2)
}
