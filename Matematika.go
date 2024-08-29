package main

import "fmt"

func main(){
	var a = 10
	var b = 10
	var c = 5
	var d = 5
	var e = 15

	var f = a*b+c/d-e
	fmt.Println("Hasil dari 10*10+5/5-15 = ", f)

	var i = 15
	i += 10
	fmt.Println("Menggunakan += ", i)
	i -= 10
	fmt.Println("Menggunakan -= ",i)
	i *= 10
	fmt.Println("Menggunakan *= ",i)
	i /= 10
	fmt.Println("Menggunakan /= ",i)
	i %= 10
	fmt.Println("Menggunakan %= ",i)

	var j = 69
	j++
	fmt.Println("Menggunakan ++ ",j)
	j--
	fmt.Println("Menggunakan -- ",j)
}