// Go is a compiled time language
// Go also handles text UTF8 natively; so any text can be used 
// to compile, link and run: $ go run first.go

// You may want to have an executable; $ go build first.go
// To run the executable: $ ./first

package main 

import "fmt"

func main() {
	fmt.Println("hello, there")
	fmt.Println("notice UTF8 compatability: 字学中文")
}