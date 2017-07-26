package main

import (
	"errors"
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

$ demo.v2 (options...) [N] [text]
`

func mainInner() error {
	versionFlag := flag.Bool("version", false, "print the version information")
	flag.Usage = func() {
		fmt.Println(strings.TrimSpace(flagUsage) + "\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *versionFlag {
		fmt.Println("Project: demo.v2 (https://github.com/AstromechZA/BuildingEffectiveCLIApps)")
		fmt.Printf("Version: %s\n", Version)
		fmt.Printf("Git Commit: %s\n", CommitSHA)
		fmt.Printf("Git Date: %s\n", CommitDate)
		os.Exit(0)
	}

	if len(os.Args) != 3 {
		return errors.New("expected 2 arguments [N] [text]; see --help for info")
	}

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return fmt.Errorf("failed to parse argument 1 as a number: %s", err)
	}
	if x < 0 {
		return fmt.Errorf("argument 1 must be a positive integer")
	}

	y := strings.TrimSpace(os.Args[2])
	if len(y) == 0 {
		return fmt.Errorf("argument 2 must be a non-empty string")
	}

	for i := 0; i < x; i++ {
		fmt.Printf("Hëllo\t%d\t%s\n", i, y)
	}
	return nil
}

func main() {
	if err := mainInner(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
