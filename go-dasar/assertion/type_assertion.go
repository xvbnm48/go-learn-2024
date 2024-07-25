package main

import "fmt"

func random() any {
	return true
}
func main() {
	//var result any = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	result2 := random()
	switch result2.(type) {
	case string:
		fmt.Println("String", result2.(string))
	case int:
		fmt.Println("Int", result2.(int))
	default:
		fmt.Println("unknown")
	}
}
