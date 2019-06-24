package ch1

import (
	"fmt"
	"os"
)

// Ex1 -
func Ex1() {
	fmt.Println(os.Args[0])
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
