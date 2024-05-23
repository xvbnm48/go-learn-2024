package main

func endApp() {
	println("end app")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("app error")
	} else {
		println("app running properly")
	}
}

func main() {
	runApp(true)
	//runApp(false)
}
