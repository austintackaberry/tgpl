package ch1

import (
	"fmt"
	"os"
	"strconv"
)

// Ex2 -
func Ex2() {
	for index, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(index) + " " + arg)
	}
}
