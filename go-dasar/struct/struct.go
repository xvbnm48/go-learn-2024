package main

import "fmt"

type customer struct {
	Name  string
	Age   int
	Email string
}

func main() {
	c1 := customer{}
	c1.Name = "vini"
	c1.Age = 20
	c1.Email = "vini@hd.com"

	allCustomer := []customer{
		{Name: "vini", Age: 20, Email: "vini@hd.com"},
		{
			Name:  "riichan",
			Age:   23,
			Email: "riichan@hd.com",
		},
		{
			Name:  "meia",
			Age:   22,
			Email: "meia@hd.com",
		},
	}
	c3 := customer{"fariz", 22, "muhamadfarizwisnuprananda@hd.com"}
	fmt.Println(c1)
	fmt.Println(c1.Name)
	fmt.Println(allCustomer)
	fmt.Println(c3)
}
