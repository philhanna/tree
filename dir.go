package tree

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"sort"
	"strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// Dir is an implementation of INode for a directory
type Dir struct {
	name     string  // Directory name
	parent   *Dir    // Containing directory
	level    int     // How far removed from root node
	children []INode // Immediate children
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewDir creates a new directory object and loads its children
func NewDir(dirname string, parent *Dir) (*Dir, error) {

	dirname = strings.TrimSuffix(dirname, "/")

	var pString string
	if parent != nil {
		pString = parent.GetName()
	}
	log.Printf("DEBUG: Entering NewDir for %s, parent=%s\n", dirname, pString)

	// Create the directory object
	dir := new(Dir)
	dir.name = dirname
	dir.parent = parent
	dir.level = 0
	if parent != nil {
		dir.level = 1 + parent.GetLevel()
	}
	dir.children = make([]INode, 0)

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
			subFile := NewFile(name, dir)
			dir.children = append(dir.children, subFile)
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

// String returns a string representation of this directory
func (p *Dir) String() string {
	parts := make([]string, 0)

	// Name
	parts = append(parts, fmt.Sprintf("Name:%q", p.GetName()))

	// Parent
	var parentName string
	parent := p.GetParent()
	switch parent {
	case nil:
		parentName = "nil"
	default:
		parentName = parent.GetName()
	}
	parts = append(parts, fmt.Sprintf("Parent:%q", parentName))

	// Level
	parts = append(parts, fmt.Sprintf("Level:%d", p.GetLevel()))

	// Children
	parts = append(parts, fmt.Sprintf("Children:%v", p.children))
	
	// Done
	return strings.Join(parts, ",")
}

// ---------------------------------------------------------------------
// Implementation of INode interface
// ---------------------------------------------------------------------

func (p *Dir) GetName() string {
	return p.name
}

func (p *Dir) GetParent() *Dir {
	return p.parent
}

func (p *Dir) GetLevel() int {
	return p.level
}

func (p *Dir) IsLast() bool {
	if p.GetParent() == nil {
		return true
	}
	siblings := p.GetParent().children
	n := len(siblings)
	return p.GetName() == siblings[n-1].GetName()
}
