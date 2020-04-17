package ch4

import (
	"bufio"
	"fmt"
	"os"
)

func dup4(files []string) []string {
	var dupFiles []string
	if len(files) == 0 {
		if dupLines(os.Stdin) {
			dupFiles = append(dupFiles, os.Stdin.Name())
		}
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
				continue
			}
			if dupLines(f) {
				dupFiles = append(dupFiles, f.Name())
			}
			f.Close()
		}
	}
	return dupFiles
}

func dupLines(f *os.File) bool {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for _, c := range counts {
		if c > 1 {
			return true
		}
	}
	return false
}
