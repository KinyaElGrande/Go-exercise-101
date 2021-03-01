package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"regexp"
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

func parseURL(linkToGet *url.URL, content string) ([]string, error){
	var(
		err error
		links []string 
		matches [][]string
		findLinks = regexp.MustCompile("<a.*?href=\"(.*?)\"")
	)

	links = make([]string, 0)
	matches = findLinks.FindAllStringSubmatch(content, -1)

	for _, val := range matches {
		var linkURL *url.URL

		if linkURL, err = url.Parse(val[1]); err != nil {
			return links, err
		}

		if linkURL.IsAbs() {
			links = append(links, linkURL.String())
		} else {
			links = append(links, linkToGet.Scheme+"://"+linkToGet.Host+linkURL.String())
		}
	}

	return links, err
}
