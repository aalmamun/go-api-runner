package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args[0] = program name
	// os.Args[1] = first argument
	if len(os.Args) < 2 {
		fmt.Println("No input provided")
		return
	}

	input := os.Args[1]
	fmt.Println("Received input:", input)
}
