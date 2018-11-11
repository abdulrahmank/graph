package root_test

import (
	"github.com/abdulrahmank/graph/root"
	"github.com/abdulrahmank/graph/root/iteration"
	"github.com/abdulrahmank/graph/root/recursion"
	"testing"
)

func TestDepthFirstSearch(t *testing.T) {

	nodes := make(map[string]root.NodeList)
	nodes["a"] = root.NodeList{"b", "c", "d"}
	nodes["b"] = root.NodeList{"e", "f"}
	nodes["c"] = root.NodeList{"g"}
	nodes["e"] = root.NodeList{"h", "i"}
	expectedResult := []string{"a","b","e", "h", "i", "f", "c", "g", "d"}

	t.Run("Should traverse given tree in depth first search using recursion", func(t *testing.T) {
		result := recursion.Dfs("a", nodes)

		for index, node := range expectedResult {
			if result[index] != node {
				t.Errorf("Expected %s but was %s", node, result[index])
			}
		}
	})

	t.Run("Should traverse given tree in depth first search using iteration", func(t *testing.T) {
		result := iteration.Dfs("a", nodes)

		for index, node := range expectedResult {
			if result[index] != node {
				t.Errorf("Expected %s but was %s", node, result[index])
			}
		}
	})

}

