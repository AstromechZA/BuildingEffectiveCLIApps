package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	x, _ := strconv.Atoi(os.Args[1])
	y := os.Args[2]
	for i := 0; i < x; i++ {
		fmt.Printf("Hello  %d\t%s\n", i, y)
	}
	fmt.Printf("Done.")
}
