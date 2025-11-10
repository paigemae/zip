package unzipper

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Extract extracts a zip archive to the current directory
func Extract(zipFile, destination string) error {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if err := extractFile(f, destination); err != nil {
			return err
		}
	}
	return nil
}

func extractFile(f *zip.File, destination string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	// prepend the path with the destination folder
	// f.Name is the full path to the file inside the zip archive
	// so we need to remove the first part of the path
	// and prepend the destination folder
	path := filepath.Join(destination, f.Name[len(filepath.Dir(f.Name))+1:])

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
