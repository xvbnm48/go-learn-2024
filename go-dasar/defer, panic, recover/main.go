package main

import "fmt"

// logging adalah function yang akan dipanggil di akhir
func logging() {
	fmt.Println("selesai memanggil function ")
}
func runApp() {
	defer logging()
	fmt.Println("run the app")
}
func main() {
	runApp()
}
