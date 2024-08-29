package main

import "fmt"

func random() any{
	return true
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	value := result.(any)
	switch value.(type){
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown type", value)
	}
}