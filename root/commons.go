package root

type NodeList []string

type Node struct {
	Value string
	Edges []Edge
}

type Edge struct {
	To *Node
}
