package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var i = 10
	i += 10
	fmt.Println(i)

	i++
	fmt.Println(i)

	i--
	i--
	fmt.Println("i now with value", i)
	strI := strconv.Itoa(i)
	fmt.Println("i now is a string", strI)
	strI = reflect.TypeOf(strI).String()
	fmt.Println("strI is a", strI)
	// Check if i is a string
	if reflect.TypeOf(strI).Kind() == reflect.String {
		fmt.Println("i is a string")
	} else {
		fmt.Println("i is not a string")
	}
}
