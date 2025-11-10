package zipper

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Create creates a zip archive with the specified sources
func Create(output string, sources []string) error {
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	for _, source := range sources {
		if err := addToZip(w, source); err != nil {
			return err
		}
	}
	return nil
}

func addToZip(w *zip.Writer, source string) error {
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