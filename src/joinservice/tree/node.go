package tree

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