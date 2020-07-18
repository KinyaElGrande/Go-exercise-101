package main

import (
	"fmt"
	"sync"
)

func countUp() {
	for n := 1; n <= 25; n++ {
		fmt.Println(n)
	}
}
func countDown() {
	for n := 25; n >= 1; n-- {
		fmt.Println(n)
	}
}

func main() {
	// Declare a wait group and set the count to two.
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Printf("Count up\n")
		countUp()

		wg.Done()
	}()

	go func() {
		fmt.Printf("Count Down\n")
		countDown()

		wg.Done()
	}()

	fmt.Println("waiting to finish")
	wg.Wait()

	fmt.Println("\n terminating the program")
}
