package main

func main() {
	// make a closure
	// closure is a function that refer to variables from outside its body
	// the function can access and assign the variables from outside its body
	counter := 0

	increment := func() {
		println("Increment")
		counter++
	}
	increment()
	increment()
	println(counter)
}
