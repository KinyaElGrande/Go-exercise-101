package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "get":
		getURLCmd := flag.NewFlagSet("get", flag.ExitOnError)
		msgFlag := getURLCmd.String("url", "Pass in the url after get command", "pass in the get --url example.com")
		getURLCmd.Parse(os.Args[2:])

		if msgFlag != nil && *msgFlag != "" {
			fmt.Printf("Processing :%s Url", *msgFlag)
		} else {
			fmt.Printf("Usage : %s", *msgFlag)
		}

	case "help":
		fmt.Println("Help messages ...")

	default:
		fmt.Printf("Unknown command:- %s\n", cmd)
	}
}
