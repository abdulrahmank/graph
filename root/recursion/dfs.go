package recursion

import commons "github.com/abdulrahmank/graph/root"

func Dfs(root string, adjacencyList map[string]commons.NodeList) commons.NodeList {
	Graph := commons.CreateGraph(adjacencyList)
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
