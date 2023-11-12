// tree is a workalike for the unix tree command
package tree

// Package-global variables
var (

	// FlagA is the value of the -a command line option. If true, then
	// the tree includes hidden files.
	FlagA bool

	// FlagD is the valud of the -d command line option. If true, then
	// only directories are printed
	FlagD bool

	// FlagL is the value of the -L command line option. If specified,
	// limits the depth of the subdirectories included.
	FlagL int

	// FlagNoReport is the value of the --noreport command line option.
	// If true, the program does not print the number of directories and
	// files at the end of the tree.
	FlagNoReport bool

	// NFiles is used to keep track of how many files have been printed.
	NFiles int

	// NDirs is used to keep track of how many directories have been
	// printed.
	NDirs int
)
