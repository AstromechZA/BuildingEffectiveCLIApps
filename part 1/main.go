package main

import (
	"flag"
	"fmt"
)

var Version = "unknown"
var CommitSHA = "unknown"
var CommitDate = ""

func main() {
	versionFlag := flag.Bool("version", false, "Print the version information")
	flag.Parse()
	if *versionFlag {
		fmt.Printf("Version: %s\n", Version)
		fmt.Printf("Commit: %s (%s)\n", CommitSHA, CommitDate)
	}
}
