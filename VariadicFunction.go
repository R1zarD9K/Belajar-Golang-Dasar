package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func main() {
	// total := sumAll(10, 20, 30, 50)
	// fmt.Println(total)
	fmt.Println(sumAll(10, 20, 30, 50, 100, 200))

	numbers := []int{10, 20, 30, 50}
	fmt.Println(sumAll(numbers...))
}
