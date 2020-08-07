# books

This set of exercises involves creating a "pipeline", a concept that requires
 channels and goroutines.
 
## Exercise 1a: A Simple Pipeline

For this exercise, you will create a simple Go pipeline.

(The concept of pipelines in Go is explained in [this classic Go blog post on
 the subject](https://blog.golang.org/pipelines). The main points of that
  article are still valid, though some of the specifics may be a little
   dated. The part of the article up to the "Stopping short" section are the
    most relevant.)
    
Take a look at the file `simple\squares.go`. This file contains a function that
 generates a random number, as well as a function that takes a number and
  squares it. Each function is bit slow at doing its work.
  
The function `NoPipleline()` uses both these functions to generate 50 squares
 and print them out. It also prints the time elapsed. (It runs when you do
  `go run main.go`)
 
**Create a new function, `Pipeline()`, that uses the pipeline pattern to
 generate and print 50 squares concurrently.**

Hint: Your pipeline will run three functions that you write, perhaps named:
* `GenerateNumbers`
* `SquareNumbers`    
* `PrintNumbers`
 
## Exercise 1b: Multiple Concurrent Workers

For this exercise, take your solution from Exercise 1a and make it run faster
 by running multiple instances (5 might be a good number) of `SquareNumbers` at  the same time.

## Exercise 2a: Overlapping Words

**Given the name of an author, write a program that prints out a list of all
 the words that appear in _all_ of their books (of four letters or more). Use
  a pipleline in order to allow for concurrent processing.**
 
----------

For example, author SANDY has written three books:

"Baseball Is Fun"
```
This is a book about baseball. It is a story of a time I found myself looking
 for answers, and discovered them in a little, sewn ball.
```

"A Day to Remember"
```
The first time I rode on the bus, I had a ball. There was even an exciting
 moment when a baseball player came on the bus.
```

"Best Sports"
```
Baseball is a sport, but soccer (also known as football) is so much better. I
 have been a soccer player for as long as I can remember. One thing in common
 between these two sports is that they each rely on a ball.
```

If we supply our program with the name "SANDY" we will get back: `baseball
, ball`

----------

To solve this problem, you have been given two services, shown here:
* `booklist.Service`
    * `.GetByAuthor(authorName string) []*Book`
* `bookdetails.Service`
    * `.Get(filename string) (io.ReadCloser, error)`

Their usage is shown in the function `tryBookFunctions()` in `main.go`. Running
`go run main.go` will cause this function to be executed. **NOTE: You must
 update `dataDirectory` to your directory in order for this to work.**

## Exercise 2b: Cancellation

In this exercise we will add the ability to cancel an in-progress pipeline. 

**Gracefully shut down your program when you receive cancellation from the
 operating system (Ctrl-C). Do this by adding a `context.Context` to each
  function in the pipeline from Exercise 2a and paying attention to it getting
   canceled.**
