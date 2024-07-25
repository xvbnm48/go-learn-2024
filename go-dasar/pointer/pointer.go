package main

import "fmt"

type idolList struct {
	name, member, nationality string
}

func main() {
	//idol := idolList{name: "hira dazzle", member: "vini", nationality: "indonesia"}
	//idol2 := idol // copy value from idol to idol2
	//idol2.name = "nogizaka46"
	//
	//fmt.Println(idol)
	//fmt.Println(idol2)
	// short hand
	// idol := idolList{name: "hira dazzle", member: "vini", nationality: "indonesia"}
	//idol2 := &idol
	var idol idolList = idolList{name: "hira dazzle", member: "vini", nationality: "indonesia"}
	var idol2 *idolList = &idol
	idol2.name = "nogizaka46"

	fmt.Println(idol) // ikut berubah
	fmt.Println(idol2)
}
