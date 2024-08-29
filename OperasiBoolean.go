package main

import "fmt"

func main(){
	var nilaiAkhir = 90
	var nilaiTugas = 80

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusNilaiTugas = nilaiTugas > 80

	// && = and
	var lulus = lulusNilaiAkhir && lulusNilaiTugas
	fmt.Println(lulus)

	// || = or
	lulus = lulusNilaiAkhir || lulusNilaiTugas
	fmt.Println(lulus)

	// ! = not
	lulus = !lulusNilaiAkhir
	fmt.Println(lulus)
}