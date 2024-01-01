package main

import "fmt"

func main() {
	var str string = "Hello, World!"
	var name = "John Doe"
	fmt.Println(name)
	fmt.Println(string(str[1])) // Get the character at index 1 (index starts from 0)
	fmt.Println(str[0:5])       // Get the first 5 characters
	fmt.Println(len(str))       // Get the length of the string
	fmt.Println(string(str[4])) // Get the character at index 5 (index starts from 0)
}
