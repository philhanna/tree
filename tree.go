package tree

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type Dir struct {
	Name     string // Directory name
	Path     string // Full path
	Children []any  // Immediate children
}

type File struct {
	Name string // File name
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// NewDir creates a new directory object and loads its children
func NewDir(dirname string, path string) (*Dir, error) {

	// Create the directory object
	dir := &Dir{
		Name:     dirname,
		Path:     path,
		Children: make([]any, 0),
	}

	// Open the directory
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	// Read the directory contents (fileinfo objects)
	files, err := fp.Readdir(-1)
	if err != nil {
		return nil, err
	}

	// Sort the contents (ignoring case)
	sort.Slice(files, func(i, j int) bool {
		iUpper := strings.ToUpper(files[i].Name())
		jUpper := strings.ToUpper(files[j].Name())
		return iUpper < jUpper
	})

	// Walk through the contents and create the children of this directory
	for i := 0; i < len(files); i++ {
		file := files[i]
		name := file.Name()

		// Ignore hidden files unless -a was specified
		if !FlagA {
			if strings.HasPrefix(name, ".") {
				continue
			}
		}

		if file.IsDir() {
			subDir, err := NewDir(name, path + "/" + name)
			if err != nil {
				return nil, err
			}
			dir.Children = append(dir.Children, subDir)
		} else {
			subFile := NewFile(name)
			dir.Children = append(dir.Children, subFile)
		}
	}

	// Normal return
	return dir, nil
}

// NewFile creates a new file entry
func NewFile(filename string) *File {
	file := &File{
		Name: filename,
	}
	return file
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// PrintTree writes the directory tree at this level to stdout
func (p *Dir) PrintTree(level int) {

	fmt.Printf("%s\n", p.Name)

	for _, child := range p.Children {
		switch v := child.(type) {
		case *Dir:
			v.PrintTree(level + 1)
		case *File:
			v.PrintTree(level + 1)
		}
	}
}

func (p *File) PrintTree(level int) {
	fmt.Printf("%s\n", p.Name)
}