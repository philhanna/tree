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
	usage = `usage: tree [-acdfghilnpqrstuvxACDFJQNSUX] [-L level [-R]] [-H  baseHREF]
	[-T title] [-o filename] [-P pattern] [-I pattern] [--gitignore]
	[--matchdirs] [--metafirst] [--ignore-case] [--nolinks] [--inodes]
	[--device] [--sort[=]<name>] [--dirsfirst] [--filesfirst]
	[--filelimit #] [--si] [--du] [--prune] [--charset X]
	[--timefmt[=]format] [--fromfile] [--noreport] [--version] [--help]
	[--] [directory ...]
  ------- Listing options -------
  -a            All files are listed.
  -d            List directories only.
  -l            Follow symbolic links like directories.
  -f            Print the full path prefix for each file.
  -x            Stay on current filesystem only.
  -L level      Descend only level directories deep.
  -R            Rerun tree when max dir level reached.
  -P pattern    List only those files that match the pattern given.
  -I pattern    Do not list files that match the given pattern.
  --gitignore   Filter by using .gitignore files.
  --ignore-case Ignore case when pattern matching.
  --matchdirs   Include directory names in -P pattern matching.
  --metafirst   Print meta-data at the beginning of each line.
  --info        Print information about files found in .info files.
  --noreport    Turn off file/directory count at end of tree listing.
  --charset X   Use charset X for terminal/HTML and indentation line output.
  --filelimit # Do not descend dirs with more than # files in them.
  -o filename   Output to file instead of stdout.
  ------- File options -------
  -q            Print non-printable characters as '?'.
  -N            Print non-printable characters as is.
  -Q            Quote filenames with double quotes.
  -p            Print the protections for each file.
  -u            Displays file owner or UID number.
  -g            Displays file group owner or GID number.
  -s            Print the size in bytes of each file.
  -h            Print the size in a more human readable way.
  --si          Like -h, but use in SI units (powers of 1000).
  -D            Print the date of last modification or (-c) status change.
  --timefmt <f> Print and format time according to the format <f>.
  -F            Appends '/', '=', '*', '@', '|' or '>' as per ls -F.
  --inodes      Print inode number of each file.
  --device      Print device ID number to which each file belongs.
  ------- Sorting options -------
  -v            Sort files alphanumerically by version.
  -t            Sort files by last modification time.
  -c            Sort files by last status change time.
  -U            Leave files unsorted.
  -r            Reverse the order of the sort.
  --dirsfirst   List directories before files (-U disables).
  --filesfirst  List files before directories (-U disables).
  --sort X      Select sort: name,version,size,mtime,ctime.
  ------- Graphics options -------
  -i            Don't print indentation lines.
  -A            Print ANSI lines graphic indentation lines.
  -S            Print with CP437 (console) graphics indentation lines.
  -n            Turn colorization off always (-C overrides).
  -C            Turn colorization on always.
  ------- XML/HTML/JSON options -------
  -X            Prints out an XML representation of the tree.
  -J            Prints out an JSON representation of the tree.
  -H baseHREF   Prints out HTML format with baseHREF as top directory.
  -T string     Replace the default HTML title and H1 header with string.
  --nolinks     Turn off hyperlinks in HTML output.
  ------- Input options -------
  --fromfile    Reads paths from files (.=stdin)
  ------- Miscellaneous options -------
  --version     Print version and exit.
  --help        Print usage and this help message and exit.
  --            Options processing terminator.
  `
)

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

func main() {

	flag.BoolVar(&tree.Flag_a, "a", false, "All files are listed.")
	flag.BoolVar(&tree.Flag_d, "d", false, "List directories only.")
	flag.IntVar(&tree.Flag_L, "L", 0, "Descend only level directories deep.")
	flag.BoolVar(&tree.Flag_noreport, "noreport", false, "Turn off file/directory count at end of tree listing.")

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

	tree.PrintTree(dir)
	if !tree.Flag_noreport {
		fmt.Printf("\n%d directories, %d files\n", tree.NDirs, tree.NFiles)
	}
}
