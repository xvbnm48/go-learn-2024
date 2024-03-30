package main

import "fmt"

func main() {
	//counter := 1
	for counterr := 1; counterr <= 10; counterr++ {
		fmt.Println("counter ke", counterr)
	}
	//for counter <= 10 {
	//	println(counter)
	//	counter++
	//}
	idolMap := map[string]string{
		"keyakizaka46": "sugai yuuka",
		"nogizaka46":   "saito asuka",
		"hkt48":        "tanaka miku",
		"akb48":        "yokoyama yui",
	}
	for key, value := range idolMap {
		fmt.Println(key, value)
	}
	idol := []string{
		"hira dazzle",
		"HKT48",
		"nogizaka46",
		"keyakizaka46",
	}
	fmt.Println(idol)
	for i := 0; i < len(idol); i++ {
		fmt.Println(idol[i])
	}
	for _, idols := range idol {
		fmt.Println(idols)
	}
	fmt.Println("counter selesai")
}
