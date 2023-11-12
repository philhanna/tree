package tree

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// File is an implementation of INode for a file
type File struct {
	name   string // File name
	parent *Dir   // Containing directory
	level  int    // How far removed from root node
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
	if p.GetParent() == nil {
		return true
	}
	siblings := p.GetParent().children
	n := len(siblings)
	return p.GetName() == siblings[n-1].GetName()
}
