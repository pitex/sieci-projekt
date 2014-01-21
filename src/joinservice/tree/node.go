package tree

import (
	"os"
	"fmt"
)

type Node struct {
	IP		  string

	// number of children
	size      int

	// all children, size = computer's capacity
	children  []*Node
}

func NewNode(ip string, capacity int) *Node {
	chld := make([]*Node, capacity)

	return &Node{ip, 0, chld}
}

// Data about node in format required by Google Charts.
func NodeFormatted(node *Node, parent string, ToolTip string) string {
	return fmt.Sprintf("['%s','%s','%s'],", node.IP, parent, ToolTip)
}

func DFS(node *Node, parent string, file *os.File) {
	file.WriteString(NodeFormatted(node, parent, ""))
	for i := 0; i < node.size; i++ {
		DFS(node.children[i], node.IP, file)
	}
}