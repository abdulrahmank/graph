package dp

import commons "github.com/abdulrahmank/graph/root"

func Dfs(root string, adjacencyList map[string]commons.NodeList) commons.NodeList {
	Graph := commons.CreateGraph(adjacencyList)
	rootNode := Graph[root]
	// Remove the root node from graph
	delete(Graph, root)
	result := commons.NodeList{root}

	// find the DFS node list for all the sub graphs [Sub Tasking]
	dfsEntries := make(map[string]commons.NodeList)
	for key := range Graph {
		dfsEntries[key] = dfs(Graph[key], dfsEntries)
	}

	//append all the results together
	for _, edge := range rootNode.Edges {
		result = append(result, dfsEntries[edge.To.Value]...)
	}

	return result
}

func dfs(rootNode *commons.Node, entries map[string]commons.NodeList) commons.NodeList {
	toBeVisited := append([]*commons.Node{rootNode}, make([]*commons.Node, 0, 0)...)
	visited := commons.NodeList{}

	for len(toBeVisited) != 0 {
		node := toBeVisited[0]
		visited = append(visited, node.Value)
		toBeVisited = toBeVisited[1:]
		tmpToBeVisited := make([]*commons.Node, 0, 0)
		for _, edge := range node.Edges {
			if entries[edge.To.Value] != nil {
				visited = append(visited, entries[edge.To.Value]...)
			} else {
				tmpToBeVisited = append(tmpToBeVisited, edge.To)
			}
		}
		toBeVisited = append(tmpToBeVisited, toBeVisited...)
	}

	return visited
}
