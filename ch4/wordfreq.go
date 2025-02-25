package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "open file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords) //将文本行按照单词分割而不是按行分割
	for input.Scan() {
		word := input.Text()
		word = strings.ToLower(word)
		counts[word]++
	}

	for word, count := range counts {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}

}
