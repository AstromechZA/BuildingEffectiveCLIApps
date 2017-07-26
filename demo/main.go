package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
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
