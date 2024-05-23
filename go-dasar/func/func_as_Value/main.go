package main

import "fmt"

func hello(name string) string {
	return "hello " + name
}

func main() {
	helloVini := hello
	fmt.Println(helloVini("vini"))
}
