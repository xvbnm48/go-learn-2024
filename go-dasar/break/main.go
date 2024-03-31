package main

func main() {
	// perulangan break
	//for i := 0; i < 10; i++ {
	//	if i == 5 {
	//		break
	//	}
	//	println(i)
	//}

	for a := 0; a < 10; a++ {
		if a%2 == 0 {
			continue
		}
		println(a)
	}
}
