package root

type NodeList []string

type Node struct {
	Value string
	Edges []Edge
}

type Edge struct {
	To *Node
}

func CreateGraph(adjacencyList map[string]NodeList) map[string]*Node {
	Graph := make(map[string]*Node)
	for key, value := range adjacencyList {
		if Graph[key] == nil {
			Graph[key] = &Node{Value: key}
		}
		if len(Graph[key].Edges) == 0 {
			Graph[key].Edges = make([]Edge, 0, 10)
		}
		for _, toNode := range value {
			if Graph[toNode] == nil {
				Graph[toNode] = &Node{Value: toNode}
			}
			Graph[key].Edges = append(Graph[key].Edges, Edge{To: Graph[toNode]})
		}
	}
	return Graph
}
