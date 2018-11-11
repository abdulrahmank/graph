package iteration

import commons "github.com/abdulrahmank/graph/root"

func Dfs(root string, adjacencyList map[string]commons.NodeList) commons.NodeList {
	Graph := commons.CreateGraph(adjacencyList)
	return dfs(Graph[root])
}

func dfs(rootNode *commons.Node) commons.NodeList {
	toBeVisited := append([]*commons.Node{rootNode}, make([]*commons.Node, 0, 0)...)
	visited := commons.NodeList{}

	for len(toBeVisited) != 0 {
		node := toBeVisited[0]
		visited = append(visited, node.Value)
		toBeVisited = toBeVisited[1:]
		tmpToBeVisited := make([]*commons.Node, 0, 0)
		for _, edge := range node.Edges {
			tmpToBeVisited = append(tmpToBeVisited, edge.To)
		}
		toBeVisited = append(tmpToBeVisited, toBeVisited...)
	}

	return visited
}
