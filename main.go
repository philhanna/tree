package main

import (
	"flag"
	"fmt"
	"os"
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
	filename := "."
	if flag.NArg() != 0 {
		filename = flag.Arg(0)
	}
	
	err := Tree(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func Tree(filename string) error {
	return nil
}