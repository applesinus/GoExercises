package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func Ex4(files []string) {
	counts := make(map[string]int)
	repeatedFiles := make(map[string]string)
	if len(files) == 0 {
		countLines(os.Stdin, counts, repeatedFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open("ch1/" + arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue

			}
			countLines(f, counts, repeatedFiles)
			_ = f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d:%v\t%s\n", n, repeatedFiles[line], line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, repeatedFiles map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		repeatedFiles[input.Text()] += " " + f.Name()[4:len(f.Name())-4]
	}
}
