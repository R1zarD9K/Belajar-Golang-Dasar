package main

import "fmt"

func main() {
	name := [...]string{"Rahmat", "Pratami", "Saputra", "Adit", "Sam", "Lutfi"}

	slice1 := name[4:6]
	fmt.Println(slice1)

	slice2 := name[:4]
	fmt.Println(slice2)	

	slice3 := name[4:]
	fmt.Println(slice3)

	slice4 := name[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daysSlice1 := days[5:]
	fmt.Println(daysSlice1)

	daysSlice1[0] = "Sabtu Baru"
	daysSlice1[1] = "Minggu Baru"
	fmt.Println(daysSlice1)
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "Libur Baru")
	fmt.Println(daysSlice2)
	daysSlice2[0] = "Sabtu Lama"
	fmt.Println(daysSlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Rahmat"
	newSlice[1] = "Pratami"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Saputra")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))

	newSlice2[0] = "Rahmat Pratami"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}

//Rata2 bentuk slice di golang jarang menggunakan array, lebih sering menggunakan slice