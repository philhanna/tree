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

// Dir is an implementation of INode for a directory
type Dir struct {
	Node
	children []INode // Immediate children
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewDir creates a new directory object and loads its children
func NewDir(dirname string, parent *Dir) (*Dir, error) {

	// Create the directory object
	dir := new(Dir)
	dir.name = dirname
	dir.parent = parent
	dir.children = make([]INode, 0)

	// Check for the -L depth valud
	if FlagL > 0 {
		if dir.GetLevel() >= FlagL {
			return dir, nil
		}
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
		if file.IsDir() {
			subDir, err := NewDir(name, dir)
			if err != nil {
				return nil, err
			}
			dir.children = append(dir.children, subDir)
		} else {
			if !FlagD {
				subFile := NewFile(name, dir)
				dir.children = append(dir.children, subFile)
			}
		}
	}

	// Normal return
	return dir, nil
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// GetPath returns the full path of this directory from the root
func (p *Dir) GetPath() string {
	switch {
	case p.parent == nil:
		return p.name
	default:
		return p.parent.GetPath() + "/" + p.name
	}
}
