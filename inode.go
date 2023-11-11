package tree

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// INode represents a file or directory
type INode interface {
	GetName() string // The node name
	GetParent() *Dir // The containing directory node
	GetLevel() int   // How far removed from root node
	IsLast() bool    // True if this is the last child of the parent
}
