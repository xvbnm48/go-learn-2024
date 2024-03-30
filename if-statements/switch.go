package main

func main() {
	var nama = "hira dazzle"

	switch nama {
	case "hira dazzle":
		println("Hello hira dazzle")
	default:
		println("Hello, World!")
	}
	switch length := len(nama); length > 5 {
	case true:
		println("Nama terlalu panjang")
	default:
		println("Nama sudah benar")
	}
}
