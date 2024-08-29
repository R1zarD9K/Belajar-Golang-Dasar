package main

import "fmt"

// Mendefinisikan interface 'HasName' yang mengharuskan adanya metode 'GetName'
type HasName interface {
	GetName() string
}

// Fungsi 'sayHello' menerima parameter 'value' yang harus mengimplementasikan interface 'HasName'
// Fungsi ini mencetak pesan "Hello" diikuti dengan nama yang diperoleh dari 'GetName'
func sayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// Mendefinisikan struct 'Person' dengan satu field 'Name'
type Person struct {
	Name string
}

// Implementasi metode 'GetName' untuk struct 'Person'
// Metode ini mengembalikan nilai dari field 'Name' pada struct 'Person'
func (person Person) GetName() string {
	return person.Name
}

// Mendefinisikan struct 'Animal' dengan satu field 'Name'
type Animal struct {
	Name string
}

// Implementasi metode 'GetName' untuk struct 'Animal'
// Metode ini mengembalikan nilai dari field 'Name' pada struct 'Animal'
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	// Membuat objek 'Person' dengan nama "Rahmat Pratami"
	person := Person{Name: "Rahmat Pratami"}
	// Memanggil fungsi 'sayHello' dengan objek 'Person'
	sayHello(person)

	// Membuat objek 'Animal' dengan nama "Gajah"
	animal := Animal{Name: "Gajah"}
	// Memanggil fungsi 'sayHello' dengan objek 'Animal'
	sayHello(animal)
}
