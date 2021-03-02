package reader

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/KinyaElGrande/Go-exercise-101/flags/samples"
)

//ReadingFlag function implementation : -
// go run main.go -n 12 main.go (prints upto line number 12 of the filename)
func ReadingFlag() {
	var count int
	flag.IntVar(&count, "n", 5, "number of lines to read from the file")
	flag.Parse()

	var in io.Reader
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error openning file : err", err)
			os.Exit(1)
		}
		defer f.Close()

		in = f
	} else {
		in = os.Stdin
	}

	buf := bufio.NewScanner(in)

	for i := 0; i < count; i++ {
		if !buf.Scan() {
			break
		}
		samples.Colorize(samples.ColorGreen, buf.Text())
	}

	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading : err:", err)
	}
}
