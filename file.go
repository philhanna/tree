package tree

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// File is an implementation of INode for a file
type File struct {
	name   string // File name
	parent *Dir   // Containing directory
	level  int    // How far removed from root node
	isLast bool   // True if this is the last child of the parent
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewFile creates a new file entry
func NewFile(filename string, parent *Dir) *File {
	file := new(File)
	file.name = filename
	file.parent = parent

	// level is one greater than parent level
	file.level = 1 + parent.GetLevel()

	// isLast is true if this file name is the same as the name of
	// the last child of the parent
	n := len(parent.children)
	lastChild := parent.children[n-1]
	nameOfLastChild := lastChild.GetName()
	file.isLast = filename == nameOfLastChild

	return file
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

func (p *File) GetName() string {
	return p.name
}

func (p *File) GetParent() *Dir {
	return p.parent
}

func (p *File) GetLevel() int {
	return p.level
}

func (p *File) IsLast() bool {
	return p.isLast
}
