package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

const (
	minWait   time.Duration = time.Millisecond * 1
	maxWait                 = time.Millisecond * 500
	minResult               = 1
	maxResult               = 1000
)

// GenerateNum returns a random integer between minResult and maxResult.
func GenerateNum() int {
	wait()
	return minResult + randomizer.Intn(maxResult-minResult)
}

// SlowSquare returns the value of in squared. It does this slowly.
func SlowSquare(in int) int {
	wait()
	return in * in
}

func wait() {
	sleepTime := minWait + time.Duration(randomizer.Int63n(int64(maxWait-minWait)))
	time.Sleep(sleepTime)
}

// NoPipeline generates 50 squares of random numbers, one after the other.
func NoPipeline() {
	fmt.Println("== Generating Squares Without a Pipeline...")
	start := time.Now()
	for i := 0; i < 50; i++ {
		in := GenerateNum()
		fmt.Printf(" * %d --> %d\n", in, SlowSquare(in))
	}
	fmt.Printf("Elapsed: %s\n", time.Now().Sub(start))
}

//SquareResult to hold channel input and output
type SquareResult struct {
	Initial int
	Square  int
}

// Generate channel
func Generate() chan int {
	numbers := make(chan int)

	go func() {
		for i := 0; i < 50; i++ {
			numbers <- GenerateNum()
		}
		close(numbers)
	}()
	return numbers
}

//Square channel
func Square(numbers chan int) chan SquareResult {
	squares := make(chan SquareResult)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for x := range numbers {
				squares <- SquareResult{Initial: x, Square: SlowSquare(x)}
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(squares)
	}()

	return squares
}

// Pipeline generates 50 squares of random numbers, concurrently.
func Pipeline() {
	fmt.Println("== Generating Squares With the Pipeline...")
	start := time.Now()

	numbers := Generate()
	squares := Square(numbers)
	//Printing the squares output
	for x := range squares {
		fmt.Printf("%+v \n", x)
	}

	fmt.Printf("Elapsed: %s\n", time.Now().Sub(start))

}

func main() {
	NoPipeline()
	// Pipeline()
}

// when can we use buffered channels ?
