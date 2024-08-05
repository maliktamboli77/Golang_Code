package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	verbose := flag.Bool("v", false, "Enable verbose mode")
	name := flag.String("name", "World", "A name to say hello to")

	// Parse the flags
	flag.Parse()

	// Use the flags
	if *verbose {
		fmt.Println("Verbose mode enabled")
	}
	fmt.Printf("Hello, %s!\n", *name)
}
