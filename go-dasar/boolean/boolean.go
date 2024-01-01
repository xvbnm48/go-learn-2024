package main

import "fmt"

func main() {
	fmt.Println("benar = ", true)
	fmt.Println("salah = ", false)
	var benar bool = true
	fmt.Println("benar = ", benar)
	var salah bool = false
	fmt.Println("salah = ", salah)

	// Example of a boolean operation
	result := benar && salah
	fmt.Println("result =", result)
	var negatif bool = false
	fmt.Println("negatif jadi positif = ", !negatif)

	// array
	var IdolJapan [3]string
	IdolJapan[0] = "Nishino Nanase"
	IdolJapan[1] = "Mai Shiraishi"
	IdolJapan[2] = "Ikuta Erika"
	fmt.Println("idol japan: ", IdolJapan[0], IdolJapan[1], IdolJapan[2])
	// slice
	var fruits = []string{"apple", "grape", "banana", "melon"}
	// add new slice
	var newFruits = append(fruits, "papaya")
	fmt.Println(newFruits)

	//make a new slice
	anime := make([]string, 2)
	anime[0] = "naruto"
	anime[1] = "one piece"

	song := []string{"naruto", "one piece", "dragon ball"}
	fmt.Println(song[1])
}
