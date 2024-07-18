package main

import "fmt"

type HahsName interface {
	GetName() string
}

func SayHello(h HahsName) {
	println("Hello", h.GetName())
}
func Idols_chika(h HahsName) {
	println("Hello", h.GetName())
}

type Idol struct {
	Name        string
	SocialMedia string
	Member      []string
}
type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

func (i Idol) GetName() string {
	return fmt.Sprintf("Name %s \n Social media %s \n Member %v \n", i.Name, i.SocialMedia, i.Member[0])

}

func main() {
	person := Person{Name: "vini"}
	idols := Idol{
		Name:        "hira dazzle",
		SocialMedia: "@hira_dazzle",
		Member: []string{
			"vini",
			"riichan",
			"meia",
			"rara",
			"eve",
			"jey",
			"jelly",
		},
	}
	Idols_chika(idols)
	SayHello(person)
}
