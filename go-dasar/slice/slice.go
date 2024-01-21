package main

import "fmt"

func main() {
	// make a slice
	var slice = []int{1, 2, 3, 4, 5}

	fmt.Println("Slice length", len(slice))
	fmt.Println("array index from low to high", slice[1:3])
	fmt.Println("array index from low to high", slice[1:])
	fmt.Println("array index high to low", slice[:3])
	fmt.Println("array index high to low", slice[:])
}
