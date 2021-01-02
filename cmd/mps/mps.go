package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kristianvalind/mini-photo-sorter/pkg/mps"
)

var showHelp, stopOnError, recurseDir, dryRun bool
var outputBasePath, outputPattern string

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] files\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&dryRun, "d", false, "dry run, simulate operations but don't move any files")
	flag.BoolVar(&showHelp, "h", false, "show usage")
	flag.StringVar(&outputBasePath, "o", ".", "`dir` in which to place output files")
	flag.StringVar(&outputPattern, "p", "2006-01-02/2006-01-02_{filename}", "the `pattern` for the renamed files, using golang time package formatting. The string {filename} is replaced with the original file name.")
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

	// Get absolute path for output
	outputBasePath, err := filepath.Abs(outputBasePath)
	if err != nil {
		fmt.Printf("could not get absolute path: %v", err)
		os.Exit(1)
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
			if recurseDir {
				err := filepath.Walk(pathToVisit, makeWalker(&filesToProcess))
				if err != nil {
					fmt.Printf("could not walk directory: %v\n", err)
					if stopOnError {
						os.Exit(1)
					} else {
						continue
					}
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

	// TODO! Implement uniqueness check

	// Process files
	for _, fileToProcess := range filesToProcess {
		fileDate, err := mps.GetDate(fileToProcess)
		if err != nil {
			fmt.Printf("could not get date for file: %v\n", err)
			if stopOnError {
				os.Exit(1)
			} else {
				continue
			}
		}

		outputFileName := strings.ReplaceAll(fileDate.Format(outputPattern), "{filename}", filepath.Base(fileToProcess))
		outputFilePath := filepath.Join(outputBasePath, outputFileName)

		fmt.Printf("%v -> %v\n", fileToProcess, outputFilePath)
	}
}
