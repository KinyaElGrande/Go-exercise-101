package main

import (
	"flag"

	"github.com/KinyaElGrande/Go-exercise-101/flags/reader"
	"github.com/KinyaElGrande/Go-exercise-101/flags/samples"
)

type GreetCommand struct {
	fs *flag.FlagSet

	name string
}

func main() {
	samples.ColoredOutput()
	reader.ReadingFlag()

}
