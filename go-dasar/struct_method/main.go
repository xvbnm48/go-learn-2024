package main

import "fmt"

type customer struct {
	Name string
	Age  int
}

type Company struct {
	Name    string
	Address string
	Email   string
	Phone   string
}

func (c customer) sayHello(name string) {
	fmt.Println("Hello ", name, " my name is ", c.Name)
}

func (C Company) greetings(name string) {
	fmt.Println("Hello ", name, " my name is ", C.Name)
}

func main() {
	c1 := customer{
		Name: "vini",
		Age:  19,
	}
	c1.sayHello("fariz")
	c2 := customer{
		Name: "meia hira dazzle",
		Age:  23,
	}

	c2.sayHello("fariz")

	// for company
	companyIndonesia := Company{
		Name:    "PT. Indonesia",
		Address: "Jl. Jendral Sudirman",
		Email:   "indonesiamaju@ikn.com",
		Phone:   "08123456789",
	}

	companyIndonesia.greetings("fariz")

}
