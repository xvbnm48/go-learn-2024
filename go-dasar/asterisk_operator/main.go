package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address := Address{
		City:     "Palembang",
		Province: "Sumatera selatan",
		Country:  "Indonesia",
	}

	address2 := &address
	address2.Province = "Jawa Barat"
	fmt.Println(address)
	fmt.Println(address2)

	//address2 = &Address{
	//	City:     "Bandung",
	//	Province: "Jawa barat",
	//	Country:  "Indonesia",
	//}
	*address2 = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}
	fmt.Println(address)
	fmt.Println(address2)
}
