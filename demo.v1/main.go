package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Version = "unknown"
var CommitSHA = "unknown"
var CommitDate = "1970-01-01T00:00:00Z"

const flagUsage = `
This binary solves the greatest programming problem known to man! It will print 'N' rows of hello text.

Usage:

$ demo (options...) [N] [text]
`

func main() {
	versionFlag := flag.Bool("version", false, "print the version information")
	flag.Usage = func() {
		fmt.Println(strings.TrimSpace(flagUsage) + "\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *versionFlag {
		fmt.Printf("Version: %s\n", Version)
		fmt.Printf("Git Commit: %s\n", CommitSHA)
		fmt.Printf("Git Date: %s\n", CommitDate)
		os.Exit(0)
	}

	if len(os.Args) == 3 {
		x, _ := strconv.Atoi(os.Args[1])
		if x < 0 {
			fmt.Printf("There was an error")
			return
		}
		y := os.Args[2]
		for i := 0; i < x; i++ {
			fmt.Printf("Hëllo  %d\t%s\n", i, y)
		}
		fmt.Printf("Done.")
	}
}
