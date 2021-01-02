package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var showHelp, stopOnError, recurseDir bool

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] files\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&showHelp, "h", false, "show usage")
	flag.BoolVar(&recurseDir, "r", false, "recurse into any provided directories")
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

	// Build a canonical list of files to process
	for _, pathToVisit := range flag.Args() {
		log.Print(pathToVisit)
	}
}
