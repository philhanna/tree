package tree

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// File is an implementation of INode for a file
type File struct {
	name   string // File name
	parent *Dir   // Containing directory
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewFile creates a new file entry
func NewFile(filename string, parent *Dir) *File {
	file := new(File)
	file.name = filename
	file.parent = parent

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
	parent := p.GetParent()
	switch parent {
	case nil:
		return 0
	default:
		return 1 + parent.GetLevel()
	}
}

func (p *File) IsLast() bool {
	if p.GetParent() == nil {
		return true
	}
	siblings := p.GetParent().children
	n := len(siblings)
	return p.GetName() == siblings[n-1].GetName()
}
