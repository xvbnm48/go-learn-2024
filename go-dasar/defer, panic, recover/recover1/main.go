package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered dari panic:", r)
			// Lakukan tindakan pemulihan atau penanganan kesalahan di sini
		}
	}()

	// Simulasi panic
	//panic("Sesuatu yang tidak terduga terjadi!")
	fmt.Println("hello, tidak ada panic")
}
