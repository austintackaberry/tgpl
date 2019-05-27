// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string](map[string]bool))
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin", filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		s := make([]string, 5)
		for filename := range filenames[line] {
			s = append(s, filename)
		}

		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, s)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string, filenames map[string](map[string]bool)) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if filenames[input.Text()] == nil {
			filenames[input.Text()] = make(map[string]bool)
		}
		filenames[input.Text()][filename] = true

	}
	// NOTE: ignoring potential errors from input.Err()
}
