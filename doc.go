// tree is a workalike for the unix tree command
package tree

// Package-global variables
var (

	// Flag_a is the value of the -a command line option. If true, then
	// the tree includes hidden files.
	Flag_a bool

	// Flag_d is the valud of the -d command line option. If true, then
	// only directories are printed
	Flag_d bool

	// Flag_L is the value of the -L command line option. If specified,
	// limits the depth of the subdirectories included.
	Flag_L int

	// Flag_noreport is the value of the --noreport command line option.
	// If true, the program does not print the number of directories and
	// files at the end of the tree.
	Flag_noreport bool

	// NFiles is used to keep track of how many files have been printed.
	NFiles int

	// NDirs is used to keep track of how many directories have been
	// printed.
	NDirs int
)
