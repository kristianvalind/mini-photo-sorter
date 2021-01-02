package main

import (
	"log"
	"os"
	"path/filepath"
)

func makeWalker(recurse bool, filesList *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		log.Print(path)
		return nil
	}
}
