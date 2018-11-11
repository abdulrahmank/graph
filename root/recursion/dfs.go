package recursion

import commons "github.com/abdulrahmank/graph/root"

func Dfs(root string, adjacencyList map[string]commons.NodeList) commons.NodeList {
	Graph := make(map[string]*commons.Node)
	for key, value := range adjacencyList {
		if Graph[key] == nil {
			Graph[key] = &commons.Node{Value: key}
		}
		if len(Graph[key].Edges) == 0 {
			Graph[key].Edges = make([]commons.Edge, 0, 10)
		}
		for _, toNode := range value {
			if Graph[toNode] == nil {
				Graph[toNode] = &commons.Node{Value: toNode}
			}
			Graph[key].Edges = append(Graph[key].Edges, commons.Edge{To: Graph[toNode]})
		}
	}

	result := commons.NodeList{}

	dfs(Graph[root], &result)

	return result
}

func dfs(rootNode *commons.Node, accumulator *commons.NodeList) {
	*accumulator = append(*accumulator, rootNode.Value)
	for _, edge := range rootNode.Edges {
		dfs(edge.To, accumulator)
	}
}
