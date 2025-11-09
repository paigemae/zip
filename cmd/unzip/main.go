package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: unzip <file.zip>")
		os.Exit(1)
	}

	zipFile := os.Args[1]

	r, err := zip.OpenReader(zipFile)
	if err != nil {
		fmt.Printf("Error opening zip file: %v\n", err)
		os.Exit(1)
	}
	defer r.Close()

	for _, f := range r.File {
		if err := extractFile(f); err != nil {
			fmt.Printf("Error extracting %s: %v\n", f.Name, err)
			os.Exit(1)
		}
	}
}

func extractFile(f *zip.File) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	path := f.Name
	if f.FileInfo().IsDir() {
		return os.MkdirAll(path, f.FileInfo().Mode())
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.FileInfo().Mode())
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, rc)
	return err
}
