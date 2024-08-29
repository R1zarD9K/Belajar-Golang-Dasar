package main

import "fmt"

//Struct adalah Prototype dari data yang akan digunakan
type Customer struct{
	Name, Address string
	Age 		  int
}

//Menggunakan methode untuk mengakses data dari struct
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

//Tapi dalam struct kita bisa membuat function
func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 30
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)
	fmt.Println(eko)

//Membuat variabel joko dari tipe Customer dengan menggunakan struct literal
	joko := Customer{
		Name: "Joko",
		Address: "Indonesia",
		Age: 35,
	}
	fmt.Println(joko)

//Membuat variabel budi dari tipe Customer dengan menggunakan struct literal tanpa menyebutkan nama field
	budi := Customer{"Budi", "Indonesia", 40}
	fmt.Println(budi)

//Menggunakan methode untuk mengakses data dari struct
	budi.sayHello("Rahmat")
}