package main

import (
	"fmt"
	"log"

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
	// wordschannel := getWords(bookchannel)
	// for _, w := range wordschannel {
	// }

	for b := range bookchannel {
		fmt.Printf("%+v \n", b)
	}
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

//book files pipeline
func bookFiles(foundBooks chan *booklist.Book) chan string {
	booksfile := make(chan string)
	detailService := bookdetails.NewService(dataDirectory)

	go func() {
		for x := range foundBooks {
			bookfile, _ := detailService.Get(x.Filename)

			defer bookfile.Close()

			start := make([]byte, 257)
			count, readErr := bookfile.Read(start)
			if readErr != nil {
				log.Fatalf("Error reading book: %s", readErr)
			}

			fmt.Printf("\nFirst %d bytes of book: %s\n", count, start)
		}

	}()

	return booksfile
}
