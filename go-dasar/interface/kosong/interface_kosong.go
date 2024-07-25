package main

import "fmt"

func Kosong() any {
	return "Hello, World!"
}

func main() {
	var kosong any = Kosong()
	fmt.Println(kosong)
}
