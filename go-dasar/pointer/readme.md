# pointer 

di golang itu semua variable itu di pass by value, jadi ketika kita mengirim variable ke function, maka yang dikirim adalah value dari variable tersebut, bukan reference dari variable tersebut.

```go
package main

import "fmt"

type idol struct {
	name string
}

func main() {
	idolList := []idol{
		{"hira dazzle"},
		{"sakura gakuin"},
		{"nogizaka46"},
	}
	idolList2 := idolList 
	idolList[0].name = "akb48"
	fmt.Println(idolList) // same data 
	fmt.Println(idolList2) // new data akb 
}
```

jadi tanpa pointer ini tuh variabel idolList2 itu adalah hasil copy dari value idolList. 


## Pengertian Pointer
pointer ini pass by reference, memiliki kemampuan untuk merefenrensikan ke lokasi data asli dari variable tersebut tanpa memduplikasi data tersebut.  Sehingga tidak boros memory karena ga banyak pakai memory. 

```go
```