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
			prefix = prefix + "│   " 
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
		for _, child := range v.children {
			PrintTree(child)
		}
	default:
	}
}
