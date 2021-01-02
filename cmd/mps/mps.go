package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var showHelp, stopOnError, recurseDir, dryRun bool
var outputPath string

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] files\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&dryRun, "d", false, "dry run, simulate operations but don't move any files")
	flag.BoolVar(&showHelp, "h", false, "show usage")
	flag.StringVar(&outputPath, "o", ".", "`dir` in which to place output files")
	flag.BoolVar(&recurseDir, "r", false, "recurse into subdirectories of provided directories")
	flag.BoolVar(&stopOnError, "s", false, "stop when encountering an error, instead of skipping to next file")
}

func main() {
	fmt.Printf("Mini Photo Sorter (mps) by Kristian Valind\n")
	flag.Parse()

	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(0)
	}

	var filesToProcess []string

	// Build a canonical list of files to process
	for _, pathToVisit := range flag.Args() {
		pathInfo, err := os.Stat(pathToVisit)
		if err != nil {
			fmt.Printf("could not stat file: %v\n", err)
			if stopOnError {
				os.Exit(1)
			} else {
				continue
			}
		}
		if pathInfo.IsDir() {
			err := filepath.Walk(pathToVisit, makeWalker(recurseDir, &filesToProcess))
			if err != nil {
				fmt.Printf("could not walk directory: %v\n", err)
				if stopOnError {
					os.Exit(1)
				} else {
					continue
				}
			}
		} else {
			absPath, err := filepath.Abs(pathToVisit)
			if err != nil {
				fmt.Printf("could not get absolute path for file: %v\n", err)
				if stopOnError {
					os.Exit(1)
				} else {
					continue
				}
			}
			filesToProcess = append(filesToProcess, absPath)
		}
	}

	log.Print(filesToProcess)
}
