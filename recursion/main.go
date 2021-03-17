package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	testCase()
}

func testCase() {
	fmt.Println("Enter the number of test cases you  would wish to have between 1 and 100 :")
	var n int

	if m, err := scanInput(&n); m != 1 {
		log.Fatal(err)
	}

	fmt.Println("Enter the intergers")
	if n < 0 || n > 100 {
		fmt.Println("Please enter a number which is less than 100 and greater than zero")
		os.Exit(0)
	}

	all := make([]int, n)
	readNumbers(all, 0, n)
	sum := sumOfSquares(all)
	fmt.Println("The sum of all the squared intergers excluding negatives is : ", sum)
}

func readNumbers(all []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := scanInput(&all[i]); m != 1 {
		log.Fatal(err)
	}
	readNumbers(all, i+1, n-1)
}

//sumOfSquares iterates through the slice and adds the squares of the integer
//values within the slice
func sumOfSquares(mylist []int) int {
	if len(mylist) <= 0 {
		return 0
	}

	//if an integer is negative it is equated to zero
	if mylist[0] <= 0 {
		mylist[0] = 0
	}
	
	curSquare := mylist[0] * mylist[0]
	return curSquare + sumOfSquares(mylist[1:])
}

func scanInput(a *int) (int, error) {
	return fmt.Scan(a)
}
