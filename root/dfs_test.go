package root_test

import (
	"github.com/abdulrahmank/graph/root"
	"github.com/abdulrahmank/graph/root/dynamic_programming"
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

	t.Run("Should traverse given tree in depth first search using Dynamic programming", func(t *testing.T) {
		result := dp.Dfs("a", nodes)

		for index, node := range expectedResult {
			if result[index] != node {
				t.Errorf("Expected %s but was %s", node, result[index])
			}
		}
	})
}

func BenchmarkDFS(b *testing.B) {

	nodes := make(map[string]root.NodeList)
	nodes["a"] = root.NodeList{"b", "c", "d"}
	nodes["b"] = root.NodeList{"e", "f"}
	nodes["c"] = root.NodeList{"g"}
	nodes["e"] = root.NodeList{"h", "i"}

	b.Run("Recursion", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			recursion.Dfs("a", nodes)
		}
	})

	b.Run("Iteration", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			iteration.Dfs("a", nodes)
		}
	})

	b.Run("Dynamic programming", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			dp.Dfs("a", nodes)
		}
	})
}

