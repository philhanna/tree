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
	Parent   *Dir   // Immediate parent directory
	Children []any  // Immediate children
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

	// Walk through the contents and create the children of this directory
	for i := 0; i < len(files); i++ {
		file := files[i]
		name := file.Name()
		if !FlagA {
			if strings.HasPrefix(name, ".") {
				continue
			}
		}
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

// PrintTree writes the directory tree at this level to stdout
//
// All but the last child of the current directory are printed like this:
//
// │   │   ├── elves.c#
//
// where the prefix consists of:
// level-1 * (vertical bar plus three spaces), followed by
// ├──<space>
//
// The last child of the current directory uses the same prefix, except
// the last part is
// └───<space>
func (p *Dir) PrintTree(level int) {

	indent := func(level int) string {
		switch level {
		case 0:
			return ""
		default:
			return strings.Repeat("│   ", level-1)
		}
	}

	fmt.Printf("%s%s\n", indent(level), p.Name)

	for i, child := range p.Children {
		elbow := "├─── "
		if i == len(p.Children)-1 {
			elbow = "└─── "
		}
		switch v := child.(type) {
		case *Dir:
			v.PrintTree(level + 1)
		case *File:
			fmt.Printf("%s%s\n", indent(level+1)+elbow, v.Name)
		default:
			fmt.Printf("BUG: Unknown type %v\n", v)
		}
	}
}
