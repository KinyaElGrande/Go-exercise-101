package main

import (
	"fmt"
	"sync"
	"time"
)

// The goal:
// We want to get a total of all the numbers in countUp
// and in coundDown... our program should print out the total.

func countUp(upchannel chan int) {
	for n := 1; n <= 25; n++ {
		time.Sleep(time.Millisecond)
		fmt.Println("UP ", n)
		upchannel <- n

	}
}
func countDown(downchannel chan int) {
	for n := 25; n >= 1; n-- {
		time.Sleep(time.Millisecond)
		fmt.Println("DOWN ", n)
		downchannel <- n

	}
}

func main() {
	var totalSum int

	ch := make(chan int, 5)
	// Declare a wait group and set the count of logical processesors to two.
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Printf("Count up\n")
		countUp(ch)

		wg.Done()
	}()

	go func() {
		fmt.Printf("Count Down\n")
		countDown(ch)

		wg.Done()
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		for num := range ch {
			totalSum = num + totalSum
			fmt.Printf("The Total sum is %d \n", totalSum)
			// totalSum += num
		}

		wg2.Done()
	}()

	fmt.Println("waiting to finish")
	wg.Wait()
	close(ch)
	wg2.Wait()

	fmt.Printf("\n terminating the program. Sum: %d\n", totalSum)
}
