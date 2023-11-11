package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/philhanna/tree"
)

const (
	usage = `usage: tree [directory]`
)

func PrintUsage() {
	fmt.Fprintln(os.Stderr, usage)
}

func main() {
	flag.Usage = PrintUsage
	flag.Parse()
	dirname := "."
	if flag.NArg() != 0 {
		dirname = flag.Arg(0)
	}
	
	err := tree.Tree(dirname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}