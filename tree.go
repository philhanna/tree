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
	Name   string  // Directory name
	Parent *Dir    // Immediate parent directory
	Children   []any  // Immediate children
}

type File struct {
	Name string // File name
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// NewDir creates a new directory object and loads its children
func NewDir(dirname string, parent *Dir) (*Dir, error) {
	
	// Create the directory object
	dir := &Dir{
		Name: dirname,
		Parent: parent,
		Children: make([]any, 0),
	}

	// Open the directory
	path := dir.GetPath()
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

	// Sort the contents
	sort.Slice(files, func(i, j int) bool {
		iUpper := strings.ToUpper(files[i].Name())
		jUpper := strings.ToUpper(files[j].Name())
		return iUpper < jUpper
	})

	// Walk through the contents and create the children of this directory
	for i := 0; i < len(files); i++ {
		file := files[i]
		name := file.Name()
		if file.IsDir() {
			subDir, err := NewDir(name, dir)
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

func (p *Dir) GetPath() string {
	if p.Parent == nil {
		return p.Name
	} else {
		return p.Parent.GetPath() + "/" + p.Name
	}
}

func (p *Dir) PrintTree(level int) {
	indent := strings.Repeat(" ", level * 4)
	fmt.Printf("%s%s\n", indent, p.Name)

}
