package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines1(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

}

// practice 1.4
func main() {
	counts := make(map[string]int)
	files := make(map[string]map[string]struct{})
	filenames := os.Args[1:]
	if len(filenames) == 0 {
		countLines1(os.Stdin, counts)
	} else {
		for _, filename := range filenames {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
				continue
			}
			input := bufio.NewScanner(file)
			for input.Scan() {
				line := input.Text()
				if files[line] == nil {
					files[line] = make(map[string]struct{})
				}
				counts[line]++
				files[line][filename] = struct{}{}
			}
			file.Close()
		}

	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, files[line])
		}
	}

}
