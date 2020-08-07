// ===================================================
// THIS IS A SOLUTION. DO NOT PEEK.
// ===================================================
package simple

import (
	"fmt"
	"sync"
	"time"
)

func GenerateNumbers() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 50; i++ {
			out <- GenerateNum()
		}
		close(out)
	}()
	return out
}

func SquareNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- SlowSquare(i)
		}
		close(out)
	}()
	return out
}

func PrintNumbers(in <-chan int) {
	fmt.Print("Squares: ")
	for i := range in {
		fmt.Printf("%d ", i)
	}
	fmt.Println("")
}

func Pipeline() {
	fmt.Println("== Generating Squares WITH a Pipeline...")
	start := time.Now()
	generated := GenerateNumbers()
	squared := SquareNumbers(generated)
	PrintNumbers(squared)
	fmt.Printf("Elapsed: %s\n", time.Now().Sub(start))
}

func ConcurrentPipeline() {
	fmt.Println("== Generating Squares WITH a CONCURRENT Pipeline...")
	start := time.Now()
	generated := GenerateNumbersConcurrently()
	squared := SquareNumbersConcurrently(generated)
	PrintNumbers(squared)
	fmt.Printf("Elapsed: %s\n", time.Now().Sub(start))
}

func GenerateNumbersConcurrently() <-chan int {
	const workercount = 10
	in := make(chan interface{}, workercount)
	out := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for inIdx := 0; inIdx < 50; inIdx++ {
			in <- nil
		}
		close(in)
		wg.Done()
	}()

	for i := 0; i < workercount; i++ {
		wg.Add(1)
		go func() {
			for _ = range in {
				out <- GenerateNum()
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func SquareNumbersConcurrently(in <-chan int) <-chan int {
	const workercount = 10
	out := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < workercount; i++ {
		go func() {
			wg.Add(1)
			for i := range in {
				out <- SlowSquare(i)
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
