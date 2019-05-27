package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func BenchmarkSlow(b *testing.B) {
	osArgs := make([]string, 3)
	osArgs[0] = "hey"
	osArgs[1] = "hello"
	osArgs[2] = "hi"
	for i := 0; i < b.N; i++ {
		slow(osArgs)
	}
}

func BenchmarkFast(b *testing.B) {
	osArgs := make([]string, 3)
	osArgs[0] = "hey"
	osArgs[1] = "hello"
	osArgs[2] = "hi"
	for i := 0; i < b.N; i++ {
		fast(osArgs)
	}
}

func slow(osArgs []string) {
	s, sep := "", ""
	for index, arg := range osArgs {
		fmt.Println(index)
		fmt.Println(arg)
		s += sep + arg
		sep = " "
	}
}

func fast(osArgs []string) {
	fmt.Println(strings.Join(osArgs, " "))
}

func main() {
	slow(os.Args[1:])
}
