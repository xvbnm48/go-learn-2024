package main

import "fmt"

func main() {
	person := map[string]string{
		"emerald-green": "vini",
		"pink":          "rara",
		"orange":        "meia",
		"blue":          "jey",
	}
	var anime map[string]string = map[string]string{}
	anime["action"] = "naruto"
	anime["slice-of-life"] = "barakamon"
	anime["comedy"] = "gintama"
	fmt.Println(person["pink"])
	fmt.Println(anime["action"])
	delete(anime, "action")
	fmt.Println(len(anime))
	fmt.Println(anime)
}
