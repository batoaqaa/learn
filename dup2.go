package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n >= 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Println(time.Now())
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	var ahmed, batoaq string = "ahmed", "batoaq"
	var i int = 0
	for input.Scan() {
		if input.Text() == "quit" {
			fmt.Printf("Hello %s %s %d\n", ahmed, batoaq, i)
			break
		}

		i++
		counts[input.Text()]++
	}
	//Note: ignoring potential errors from input.Errr()
}
