package main

import (
	"flag"
	"fmt"
	"os"
)

// ---------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------

const (
	usage = `usage: tree [directory]
`
)
func PrintUsage() {
	fmt.Fprint(os.Stderr, usage)
}

func main() {
	flag.Usage = PrintUsage
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Printf("directory is a required parameter\n")
		os.Exit(1)
	}
}