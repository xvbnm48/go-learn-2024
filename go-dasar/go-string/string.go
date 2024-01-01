package main

import "fmt"

func main() {
	// string literal with escape characters
	fmt.Println("Hello\tWorld") // Output: Hello    World
	fmt.Println("Hello\nWorld") // Output: Hello
	//         World
	fmt.Println("Hello\"World")       // Output: Hello"World
	fmt.Println("C:\\path\\to\\file") // Output: C:\path\to\file

	fmt.Println("\u65E5")
	fmt.Print("\u3042\u308a\u304c\u305f\u3046!\n")
	fmt.Println("ja" + "pan")
	idol := "nijigasaki"
	fmt.Println(string(idol[3:])) // Output: n
	for i := 0; i < len(idol); i++ {
		fmt.Printf("%q\n", idol[i]) // Output: 'n' 'i' 'j' 'i' 'g' 'a' 's' 'a' 'k' 'i'
	}
}
