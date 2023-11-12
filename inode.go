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
	var prefix string
	if node.GetLevel() > 0 {
		switch node.IsLast() {
		case false:
			prefix = "├───"
		case true:
			prefix = "└───"
		}
	}
	work := node
	for work.GetParent() != nil {
		parent := work.GetParent()
		switch parent.IsLast() {
		case false:
			prefix = "│   " + prefix
		case true:
			prefix = "    " + prefix
		}
		work = parent
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