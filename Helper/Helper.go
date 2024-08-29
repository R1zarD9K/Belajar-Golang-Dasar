package Helper

var version = "1.0.0" //Hanya bisa diakses dalam package Helper -> huruf kecil
var Application = "Belajar Golang" //Bisa diakses diluar package Helper -> huruf besar

func SayHello(name string) string {
	return "Hello " + name
}
