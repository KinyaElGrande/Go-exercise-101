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
			tmails := strings.Trim(scanner.Text(), ",")
			mails := strings.Split(tmails, ",")
			// tmails := trimQuote(mails)
			for _, n := range mails {
				if strings.Contains(n, "@") {
					trimmed := trimQuote(n)
					lines = append(lines, trimmed)
				}
			}
		}
	}
	return lines, scanner.Err()
}

//trimming quotes
func trimQuote(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// for removing duplicate emails
func unique(line []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range line {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
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
	lines, err := readLines("emails.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range lines {
		fmt.Println(i, line)
	}

	if err := writeLines(lines, "out.csv"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}
