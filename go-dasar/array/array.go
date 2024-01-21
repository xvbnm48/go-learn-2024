package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	var arr2 = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Array index %d has value %d\n", i, arr[i])
	}
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("Array index %d has value %d\n", i, arr2[i])
	}
	var arr3 = [...]int{
		1,
		2,
		3,
	}

	for i := 0; i < len(arr3); i++ {
		fmt.Println("Array index", i, "has value", arr3[i])
	}
}
