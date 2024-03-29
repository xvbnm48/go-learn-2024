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
	nama := []string{"andi", "budi", "caca"}
	result := len(nama)
	fmt.Println("result:", result)

	idol := []string{"vini", "meia", "riichan", "rara"}
	Oshi := idol[0:1] // ["vini"]
	memberHD := idol[0:4]
	fmt.Println(memberHD)
	fmt.Println(Oshi)
	fmt.Println(cap(idol)) // 4
	fmt.Println(len(memberHD))
	newIdol := make([]string, 2, 5)
	newIdol[0] = "meia"
	newIdol[1] = "riichan"
	newIdol2 := append(newIdol, "eve") // ["meia", "riichan", "eve"]
	fmt.Println(newIdol2)
	//idol[1:4] // ["meia", "riichan", "rara"]
}
