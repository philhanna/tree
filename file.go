package tree

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// File is an implementation of INode for a file
type File struct {
	Node
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
