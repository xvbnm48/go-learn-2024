package main

func main() {
	var name = "hira dazzle"
	if name == "hira dazzle" {
		println("Hello hira dazzle")
	} else if name == "vini emerald green" {
		println("Hello vini emerald green")
	} else {
		println("Hello world!")
	}
	oshi := "vini"
	if panjang := len(oshi); panjang > 5 {
		println("Nama terlalu panjang")
	} else {
		println("Nama sudah benar")
	}
}
