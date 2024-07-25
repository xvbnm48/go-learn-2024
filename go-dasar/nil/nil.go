package main

import "fmt"

//	func contoh(name string) string {
//		if name == "" {
//			return nil
//		}
//	}
func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	data := newMap("vini")
	if data == nil {
		fmt.Println("data masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
