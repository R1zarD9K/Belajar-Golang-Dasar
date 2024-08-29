package main

func main() {
	name := "Rahmat"

	switch name {
		case "Eko":
			println("Halo Eko")
		case "Joko":
			println("Halo Joko")
		default:
			println("Kenalan dong")
	}

	switch length := len(name); length > 5 {
		case true:
			println("Nama terlalu panjang")
		case false:
			println("Nama sudah benar")
	}

	length := len(name)
	switch {
		case length > 10:
			println("Nama terlalu panjang")
		case length > 5:
			println("Nama lumayan panjang")
		default:
			println("Nama sudah benar")
	}
}