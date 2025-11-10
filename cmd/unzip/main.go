package main

import (
	"fmt"
	"os"

	"github.com/paigemae/zip/pkg/unzipper"
)

func main() {
	zipFile, destination, err := getOptions()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := unzipper.Extract(zipFile, destination); err != nil {
		fmt.Printf("Error extracting zip file: %v\n", err)
		os.Exit(1)
	}
}

func getOptions() (zipFile, destination string, err error) {
	switch len(os.Args) {
	case 2:
		return os.Args[1], ".", nil
	case 3:
		return os.Args[1], os.Args[2], nil
	default:
		return "", "", fmt.Errorf("usage: unzip <file.zip> [destination]")
	}
}
