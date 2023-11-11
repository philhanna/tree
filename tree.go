package tree

import (
	"io/fs"
	"os"
	"sort"
	"strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type Dir struct {
	Name     string // Directory name
	Parent   *Dir   // Containing directory
	Children []any  // Immediate children
}

type File struct {
	Name   string // File name
	Parent *Dir   // Containing directory
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// NewDir creates a new directory object and loads its children
func NewDir(dirname string, parent *Dir) (*Dir, error) {

	// Create the directory object
	dir := &Dir{
		Name:     dirname,
		Parent:   parent,
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

	// Sort the contents (ignoring case)
	sort.Slice(files, func(i, j int) bool {
		iUpper := strings.ToUpper(files[i].Name())
		jUpper := strings.ToUpper(files[j].Name())
		return iUpper < jUpper
	})

	// Ignore hidden files unless -a was specified
	if !FlagA {
		newFiles := make([]fs.FileInfo, 0)
		for _, file := range files {
			if !strings.HasPrefix(file.Name(), ".") {
				newFiles = append(newFiles, file)
			}
		}
		files = newFiles
	}

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
			subFile := NewFile(name, dir)
			dir.Children = append(dir.Children, subFile)
		}
	}

	// Normal return
	return dir, nil
}

// NewFile creates a new file entry
func NewFile(filename string, parent *Dir) *File {
	file := &File{
		Name:   filename,
		Parent: parent,
	}
	return file
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

func (p *Dir) Print() {

}

func (p *Dir) GetPath() string {
	switch {
	case p.Parent == nil:
		return ""
	default:
		return p.Parent.GetPath() + "/" + p.Name
	}
}
