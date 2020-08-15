package main

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"

	"./bookdetails"
	"./booklist"
	// "./simple"
)

// ***CHANGE THIS TO YOUR DIRECTORY***
const dataDirectory = "./data"

func main() {
	// simple.NoPipleline()

	// These are solutions. Uncomment to see them run.
	// simple.Pipeline()
	// simple.ConcurrentPipeline()

	// Uncomment this to see the book functions in action.
	// tryBookFunctions()
	author := "William Shakespeare"
	overlappingWords(author)
}

func tryBookFunctions() {

	fmt.Println("")
	fmt.Println("Trying out book functions...")
	fmt.Println("")

	listService, err := booklist.NewService(dataDirectory)
	if err != nil {
		log.Fatalf("Unable to create list service: %s", err)
	}

	detailService := bookdetails.NewService(dataDirectory)

	authorName := "William Shakespeare"

	fmt.Println("")
	books := listService.GetByAuthor(authorName)
	fmt.Printf("Books by '%s'...\n", authorName)
	for idx, b := range books {
		fmt.Printf("* %d. Author: %s. Title: %s\n", idx, b.Author, b.Title)
	}

	if len(books) > 0 {

		firstBook, err := detailService.Get(books[0].Filename)
		if err != nil {
			log.Fatalf("Error getting book: %s", err)
		}

		defer firstBook.Close()

		start := make([]byte, 257)
		count, readErr := firstBook.Read(start)
		if readErr != nil {
			log.Fatalf("Error reading book: %s", readErr)
		}

		fmt.Printf("\nFirst %d bytes of book: %s\n", count, start)
		fmt.Println("")
	}

}

func overlappingWords(author string) {

	bookchannel := getBooks(author)
	wordschannel := bookFiles(bookchannel)

	var commonWords map[string]int
	for b := range wordschannel {
		if commonWords == nil {
			commonWords = b
		} else {
			// find words in common...
			commonWords = findCommon(commonWords, b)
		}

		// fmt.Printf("%+v \n", b)
	}

	fmt.Printf("Common words: %v\n", commonWords)
}

func findCommon(start map[string]int, newWords map[string]int) map[string]int {
	return nil
}

//pipeline for getting books
func getBooks(author string) chan *booklist.Book {

	foundBooks := make(chan *booklist.Book)
	listService, err := booklist.NewService(dataDirectory)
	if err != nil {
		log.Fatalf("Unable to create list service: %s", err)
	}

	go func() {
		for _, b := range listService.GetByAuthor(author) {
			foundBooks <- b
		}
		close(foundBooks)
	}()

	return foundBooks
}

/*
book ---> "The cat in the hat came back. The cat ate food."
our map...
the: 2
cat: 2
in: 1
came: 1
back: 1
ate: 1
food: 1
*/

//bookFiles returns the words from the foundbooks
func bookFiles(foundBooks chan *booklist.Book) chan map[string]int {
	wordschan := make(chan map[string]int)
	detailService := bookdetails.NewService(dataDirectory)

	go func() {
		for x := range foundBooks {
			wordschan <- booktoMap(x, detailService)
		}
		close(wordschan)
	}()

	return wordschan
}

func booktoMap(book *booklist.Book, detailService *bookdetails.Service) map[string]int {
	words := make(map[string]int)
	bookfile, _ := detailService.Get(book.Filename)

	defer bookfile.Close()

	scanner := bufio.NewScanner(bookfile)
	for scanner.Scan() {
		addWords(scanner.Text(), words)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}

func addWords(text string, wordMap map[string]int) {
	// add all the words in `text` to `wordMap`
	words := strings.Split(text, " ")

	// remove punctuation
	// lower-case
	// 4 letters or longer

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}

	for _, w := range words {
		processedWord := reg.ReplaceAllString(w, "")
		value, found := wordMap[processedWord]
		if found {
			wordMap[processedWord] = value + 1
		} else {
			wordMap[processedWord] = 1
		}
	}
}
