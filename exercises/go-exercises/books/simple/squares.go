package simple

import (
	"fmt"
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

const (
	minWait   time.Duration = time.Millisecond * 1
	maxWait   time.Duration = time.Millisecond * 500
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
func NoPipleline() {
	fmt.Println("== Generating Squares Without a Pipeline...")
	start := time.Now()
	for i := 0; i < 50; i++ {
		in := GenerateNum()
		fmt.Printf("%d --> %d. ", in, SlowSquare(in))
	}
	fmt.Println("")
	fmt.Printf("Elapsed: %s\n", time.Now().Sub(start))
}
