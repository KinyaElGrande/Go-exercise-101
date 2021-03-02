package samples

import (
	"flag"
	"fmt"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

//Colorize function changes the color of the text output
func Colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

//ColoredOutput function displays the message in the color provided 
//on the terminal
func ColoredOutput() {
	useColor := flag.Bool("color", false, "Display colorized output")
	flag.Parse()

	if *useColor {
		Colorize(ColorBlue, "Hello Gopher Kinya!")
		return
	}

	fmt.Println("This is a necessity")
}
