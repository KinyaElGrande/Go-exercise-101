package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "@") {
			mails := strings.Trim(scanner.Text(), "<>;(),")
			// tmails := trimQuote(mails)
			lines = append(lines, mails)
		}
	}
	return lines, scanner.Err()
}

//trimming quotes
func trimQuote(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-2]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-2]
		}
	}
	return s
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {
	lines, err := readLines("live.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range lines {
		fmt.Println(i, line)
	}

	if err := writeLines(lines, "live_mails.csv"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}
