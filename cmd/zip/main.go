package main

import (
	"fmt"
	"os"

	"github.com/paigemae/zip/pkg/zipper"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: zip <output.zip> <file/dir> [file/dir...]")
		os.Exit(1)
	}

	output := os.Args[1]
	sources := os.Args[2:]

	if err := zipper.Create(output, sources); err != nil {
		fmt.Printf("Error creating zip file: %v\n", err)
		os.Exit(1)
	}
}
