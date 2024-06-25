package main

import "fmt"

//	func endApp() {
//		println("end app")
//	}
func endAppYangBenar() {
	fmt.Println("end app")
	message := recover()
	fmt.Println("panic with message: ", message)
}

// ini cara salah menggunakan recover
//
//	func runApp(error bool) {
//		defer endApp()
//		if error {
//			panic("app error")
//		}
//		message := recover()
//		fmt.Println("panic with message: ", message)
//	}
func runAppYangBenar(error bool) {
	defer endAppYangBenar()
	if error {
		panic("app error")
	}
}

func main() {
	//runApp(true)
	//runApp(false)
	if runAppYangBenar(false); recover() != nil {
		fmt.Println("ada panic, tapi tetap jalan karena ada recover")
	} else {
		fmt.Println("tidak ada panic")
	}
}
