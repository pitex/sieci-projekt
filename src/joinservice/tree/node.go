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

// DFS writing into given file, for need of Google Charts
func DFS(node *Node, parent string, file *os.File) {
	file.WriteString(NodeFormatted(node, parent, ""))
	for i := 0; i < node.size; i++ {
		DFS(node.children[i], node.IP, file)
	}
}

// Ads new child to father Node
func AddNewChild(father *Node, child *Node) {
	father.children[father.size] = child
	father.size++
}

// Finds father for newly added server
func FindSolution(node *Node, fatherDepth int) (*Node, int) {
	l := len(node.children)
	if node.size < l {
		return node, fatherDepth + 1
	}
	resNode := node
	depths := 2000000000
	for i := 0; i < l; i++ {
		nd, d := FindSolution(node.children[i], fatherDepth + 1)
		if d < depths {
			depths = d
			resNode = nd
		}
	}
	return resNode, depths
}