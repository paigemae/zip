package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: zip <output.zip> <file/dir> [file/dir...]")
		os.Exit(1)
	}

	output := os.Args[1]
	sources := os.Args[2:]

	file, err := os.Create(output)
	if err != nil {
		fmt.Printf("Error creating zip file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	for _, source := range sources {
		if err := addToZip(w, source, ""); err != nil {
			fmt.Printf("Error adding %s: %v\n", source, err)
			os.Exit(1)
		}
	}
}

func addToZip(w *zip.Writer, source, baseDir string) error {
	info, err := os.Stat(source)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			
			relPath, _ := filepath.Rel(filepath.Dir(source), path)
			return addFileToZip(w, path, relPath)
		})
	}
	
	return addFileToZip(w, source, filepath.Base(source))
}

func addFileToZip(w *zip.Writer, filePath, zipPath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	f, err := w.Create(zipPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, file)
	return err
}
