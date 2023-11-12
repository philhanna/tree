package tree

import "fmt"

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

// AbstractNode is the base class from which File and Dir inherit
type AbstractNode struct {
	INode
	name   string // the node name
	parent *Dir   // the containing directory node
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// TreeString creates a string representation of this node
func TreeString(node INode) string {

	n := node.GetLevel()
	if n == 0 {
		return node.GetName()
	}

	// Get a list of all the parent nodes plus this node
	comps := make([]INode, n)
	comp := node
	for i := n - 1; i >= 0; i-- {
		comps[i] = comp
		comp = comp.GetParent()
	}

	// Construct the prefix

	var prefix string
	for i := 0; i < n; i++ {
		comp = comps[i]
		switch {
		case i < n-1 && !comp.IsLast():
			prefix = prefix + "│   "
		case i < n-1 && comp.IsLast():
			prefix = prefix + "    "
		case i == n-1 && !comp.IsLast():
			prefix = prefix + "├── "
		case i == n-1 && comp.IsLast():
			prefix = prefix + "└── "
		}
	}
	line := prefix + node.GetName()
	return line
}

// PrintTree walks the tree from this node down, printing each line
func PrintTree(node INode) {
	fmt.Println(TreeString(node))
	switch v := node.(type) {
	case *Dir:
		NDirs++
		for _, child := range v.children {
			PrintTree(child)
		}
	default:
		NFiles++
	}
}

func (p *AbstractNode) GetName() string {
	return p.name
}

func (p *AbstractNode) GetParent() *Dir {
	return p.parent
}

func (p *AbstractNode) GetLevel() int {
	parent := p.GetParent()
	switch parent {
	case nil:
		return 0
	default:
		return 1 + parent.GetLevel()
	}
}

func (p *AbstractNode) IsLast() bool {
	if p.GetParent() == nil {
		return true
	}
	siblings := p.GetParent().children
	n := len(siblings)
	return p.GetName() == siblings[n-1].GetName()
}
