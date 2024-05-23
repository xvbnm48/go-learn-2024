package main

// function return
func getHello(name string) string {
	return "Hello " + name
}
func pertambahan(a, b int) int {
	return a + b
}

// function returning multiple values
func getHelloAndAge(name string, age int) (string, int) {
	return "Hello " + name, age
}

// function named return value
func Idol(name string) (myName string) {
	myName = name
	return
}

// function variadic
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	arrayFunc := sumAll(10, 10, 10, 10, 10)
	println(arrayFunc)
	MyIdol := Idol("Hira Dazzle")
	println(MyIdol)
	result := pertambahan(10, 10)
	myName, age := getHelloAndAge("John Doe", 20)
	println(result)
	println(myName, "and your age is ", age)
	println(getHello("John Doe"))
	println(pertambahan(10, 10))
}
