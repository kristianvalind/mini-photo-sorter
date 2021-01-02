package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func makeWalker(filesList *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		// Only work on files, ignore directories
		if !info.IsDir() {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return fmt.Errorf("could not get absolute path for file: %w", err)
			}

			// Skip hidden files
			if !strings.HasPrefix(filepath.Base(absPath), ".") {
				*filesList = append(*filesList, absPath)
			}
		}

		return nil
	}
}
