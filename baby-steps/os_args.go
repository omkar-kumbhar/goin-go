// print all command line arguements

package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	start := time.Now()
	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// 	fmt.Println(i,os.Args[i])
	// }
	fmt.Println(strings.Join(os.Args, " "))

	fmt.Println("Name of the command that invoked it", os.Args[0])
	fmt.Println("Time taken", time.Since(start).Seconds())
}
