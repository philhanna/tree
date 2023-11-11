package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/philhanna/tree"
)

// ---------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------

const (
	usage = `usage: tree [directory]`
)

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
	}
	flag.Parse()

	// Get the directory name (default is current directory)
	dirname := "."
	if flag.NArg() != 0 {
		dirname = flag.Arg(0)
	}

	dir, err := tree.NewDir(dirname, nil)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	_ = dir

	dir.PrintTree(0)
}
