package tree

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// Node is the base class from which File and Dir inherit
type Node struct {
	INode
	name   string // the node name
	parent *Dir   // the containing directory node
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// GetName returns the file or directory name
func (p *Node) GetName() string {
	return p.name
}

// GetParent returns a pointer to the parent directory
func (p *Node) GetParent() *Dir {
	return p.parent
}

// GetLevel returns the number of parent directories, all the way back
// to root
func (p *Node) GetLevel() int {
	parent := p.GetParent()
	switch parent {
	case nil:
		return 0
	default:
		return 1 + parent.GetLevel()
	}
}

// IsLast returns true if this file or directory is the last child of
// its parent directory
func (p *Node) IsLast() bool {
	if p.GetParent() == nil {
		return true
	}
	siblings := p.GetParent().children
	n := len(siblings)
	return p.GetName() == siblings[n-1].GetName()
}
