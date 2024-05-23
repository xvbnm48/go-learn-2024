package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func sumAllNotVariadic(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	arrayFunc := sumAll(10, 10, 10, 10, 10)
	println(arrayFunc)

	arrayNotVariadic := sumAllNotVariadic([]int{10, 10, 10, 10, 10})
	fmt.Println(arrayNotVariadic)
}
